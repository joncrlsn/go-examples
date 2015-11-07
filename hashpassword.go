//
// hashpassword.go wraps the golang bcrypt password hashing library
// with base64 encoding so we can deal in strings
//
// bcrypt is based on this paper:
// https://www.usenix.org/legacy/event/usenix99/provos/provos.pdf
//
package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretPassword = "hello, my name is inigo montoya"
)

func main() {
	// Hash the real password
	hashBase64, _ := HashPassword(secretPassword)
	fmt.Println("hashBase64:", hashBase64)

	// Compare with an invalid password
	fmt.Printf("Compare should be false: %t\n", ComparePassword(hashBase64, "some invalid password"))

	// Compare with the right password
	fmt.Printf("Compare should be true: %t\n", ComparePassword(hashBase64, secretPassword))
}

// HashPassword hashes the clear-text password and encodes it as base64,
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10 /*cost*/)
	if err != nil {
		return "", err
	}

	// Encode the entire thing as base64 and return
	hashBase64 := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashBase64, nil
}

// ComparePassword hashes the test password and then compares
// the two hashes.
func ComparePassword(hashBase64, testPassword string) bool {

	// Decode the real hashed and salted password so we can
	// split out the salt
	hashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		fmt.Println("Error, we were given invalid base64 string", err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(testPassword))
	return err == nil
}
