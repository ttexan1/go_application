package domain

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PtrString returns the pointer of string
func PtrString(s string) *string {
	return &s
}

// PtrInt returns the pointer of integer
func PtrInt(i int) *int {
	return &i
}

// PtrTime returns the pointer of time.Time
func PtrTime(t time.Time) *time.Time {
	return &t
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
