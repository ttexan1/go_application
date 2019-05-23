package domain

import jwt "github.com/dgrijalva/jwt-go"

// JWTSigningKey is the secret passed in by the ENV
var JWTSigningKey = ""

// JWTClaims represents the decrypted auth token schema
type JWTClaims struct {
	WriterID int
	jwt.StandardClaims
}
