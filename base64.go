package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "Hi Mom"
	fmt.Printf("String: '%s'\n", str)

	// Encode
	encoded := base64Encode([]byte(str))
	fmt.Println("Encoded: ", encoded)

	// Decode
	decoded, ok := base64Decode(encoded)
	fmt.Printf("Decoded: '%s'   ok:%t\n", string(decoded), ok)
}

func base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func base64Decode(str string) ([]byte, bool) {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return []byte{}, false
	}
	return bytes, true
}
