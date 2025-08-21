package utils

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashpass), nil
}
func ComparePassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}


// Generate9Digit returns a 9-digit numeric string, first digit 1-9.
func Generate9Digit() (string, error) {
	const n = 9
	b := make([]byte, n)
	_,err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("rand.Read: %w", err)
	}
	
	out := make([]byte, n)
	out[0] = '1' + (b[0] % 9)      // 1..9
	for i := 1; i < n; i++ {       // 0..9
		out[i] = '0' + (b[i] % 10)
	}
	return string(out), nil
}
