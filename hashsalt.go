//
// hashsalt provides methods to hash a password with a random salt and then
// put the two together so the hash-salt can be compared with an entered password .
//
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

const (
	SALT_SIZE      = 32
	secretPassword = "hello, my name is inigo montoya"
)

func main() {
	// Hash and salt the real password
	hashSaltBase64, _ := HashAndSaltPassword(secretPassword)
	fmt.Println("hashSaltBase64:", hashSaltBase64)

	// Compare with an invalid password
	fmt.Printf("Compare should be false: %t\n", ComparePassword("some invalid password", hashSaltBase64))

	// Compare with the right password
	fmt.Printf("Compare should be true: %t\n", ComparePassword(secretPassword, hashSaltBase64))
}

// makeSalt generates a random salt
func makeSalt() ([]byte, error) {
	salt := make([]byte, SALT_SIZE, SALT_SIZE)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// HashAndSaltPassword generates a random salt, then hashes the password
// and the salt together, then appends the salt to the end of the hash and
// returns both as a base64 encoded string.
func HashAndSaltPassword(password string) (string, error) {
	salt, err := makeSalt()
	if err != nil {
		return "", err
	}

	hashSaltBase64 := HashPassword(password, salt)
	return hashSaltBase64, nil
}

// HashPassword hashes the clear-text password and the given salt together,
// then appends the salt to the end of the hash and returns both as a
// base64 encoded string.
func HashPassword(password string, salt []byte) string {
	hash := sha256.New()

	// Hash both the password and the salt together
	hash.Write([]byte(password))
	hash.Write(salt)
	hashBytes := hash.Sum(nil)

	// Append the salt bytes to the end of the hashed bytes
	hashSaltBytes := append(hashBytes, salt...)

	// Encode the entire thing as base64 and return
	hashSaltBase64 := base64.StdEncoding.EncodeToString(hashSaltBytes)

	return hashSaltBase64
}

// ComparePassword hashes the test password with the same salt used for the real
// password and then compares the two hashes.
func ComparePassword(testPassword string, hashSaltBase64 string) bool {

	// Decode the real hashed and salted password so we can
	// split out the salt
	bytes, err := base64.StdEncoding.DecodeString(hashSaltBase64)
	if err != nil {
		fmt.Println("Error, given invalid base64 string", err)
		os.Exit(1)
	}

	// Split out the salt
	salt := bytes[len(bytes)-SALT_SIZE:]

	// Hash the test password with the same salt used to hash the real password
	testHashSaltBase64 := HashPassword(testPassword, salt)

	// Return true if the hashes match
	return hashSaltBase64 == testHashSaltBase64
}
