package domain

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// Writer describes a writer
type Writer struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Email             string  `json:"email" gorm:"unique_index; not null" valid:"required,email"`
	EncryptedPassword string  `json:"-" gorm:"not null"`
	Name              string  `json:"name" gorm:"not null" valid:"required"`
	Memo              *string `json:"memo"`
	Status            string  `json:"status" gorm:"not null" valid:"in(valid|deleted)"`
}

// Writer status
const (
	WriterStatusValid   = "valid"
	WriterStatusDeleted = "deleted"
)

// SetPassword sets the password
func (w *Writer) SetPassword(pass string) *Error {
	if len(pass) < 6 {
		return NewError(http.StatusUnprocessableEntity, "password should be greater or equal than 6 characters")
	}
	w.EncryptedPassword = Encrypt(pass)
	return nil
}

// BeforeSave is called before it is saved to the database
func (w *Writer) BeforeSave() error {
	if w.EncryptedPassword == "" {
		return errors.New("password is required")
	}
	return nil
}

// CreateJWTToken creates a an encrypted token containing JWTClaims using the signing key.
func (w *Writer) CreateJWTToken(expiration time.Time) (string, *Error) {
	if JWTSigningKey == "" {
		return "", NewError(http.StatusInternalServerError, "SIGNING_KEY is empty")
	}
	claims := JWTClaims{
		WriterID: w.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signedString, err := token.SignedString([]byte(JWTSigningKey))
	if err != nil {
		return "", NewError(http.StatusInternalServerError, "Failed to sign JWT")
	}
	return signedString, nil
}

// DecryptJWTToken decrypts a token and returns the associated JWTClaims
func DecryptJWTToken(tokenString string) (*JWTClaims, *Error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(JWTSigningKey), nil
	})
	if err != nil {
		return nil, NewError(http.StatusUnauthorized, err.Error())
	}

	var ourClaims JWTClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapstructure.Decode(claims, &ourClaims)
	} else {
		return nil, NewError(http.StatusUnauthorized, "Failed to decode token")
	}
	return &ourClaims, nil
}
