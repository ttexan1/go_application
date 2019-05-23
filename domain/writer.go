package domain

import "time"

// Writer describes a writer
type Writer struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Email             string  `json:"email" gorm:"unique_index; not null" valid:"required,email"`
	EncryptedPassword string  `json:"-" gorm:"not null"`
	Name              string  `json:"name" gorm:"not null" valid:"required"`
	Memo              *string `json:"memo"`
	Status            string  `json:"status" gorm:"not null" valid:"in(temp|valid|deleted)"`
}

// Writer status
const (
	WriterStatusTemp    = "temp"
	WriterStatusValid   = "valid"
	WriterStatusDeleted = "deleted"
)
