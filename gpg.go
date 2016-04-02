package main

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"log"
	"os"
)

// create gpg keys with
// $ gpg --gen-key
// ensure you specify correct paths and passphrase

const (
	mySecretString     = "this is so very secret!"
	prefix, passphrase = "/home/jon/", "xxxxxxxxx"
	secretKeyring      = prefix + ".gnupg/secring.gpg"
	publicKeyring      = prefix + ".gnupg/pubring.gpg"
)

func main() {

	bytes, err := ioutil.ReadFile(prefix + ".password-store/Jon/facebook.com:jon_carlson@writeme.com.gpg")
	if err != nil {
		log.Fatal(err)
	}
	decBytes, err := DecryptBytes(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("decrypted stuff", string(decBytes))

	encStr, err := EncryptBytes(mySecretString)
	if err != nil {
		log.Fatal(err)
	}
	decStr, err := DecryptBase64ToString(encStr)
	if err != nil {
		log.Fatal(err)
	}
	// should be done
	log.Println("Decrypted Secret:", decStr)
}

func EncryptBytes(origBytes []byte, pubKeyringFile string) ([]byte, error) {

	// Read in public key
	keyringFileBuffer, _ := os.Open(pubKeyringFile)
	defer keyringFileBuffer.Close()
	entityList, err := openpgp.ReadKeyRing(keyringFileBuffer)
	if err != nil {
		return "", err
	}

	// encrypt string
	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entityList, nil, nil, nil)
	if err != nil {
		return "", err
	}
	_, err = w.Write(origBytes)
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}
	encryptedBytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return "", err
	}

	return encryptedBytes, nil
}

func DecryptBytes(encBytes []byte, privKeyringFile string) ([]byte, error) {

	// init some vars
	var entity *openpgp.Entity
	var entityList openpgp.EntityList

	// Open the private key file
	keyringFileBuffer, err := os.Open(privKeyringFile)
	if err != nil {
		return nil, err
	}
	defer keyringFileBuffer.Close()
	entityList, err = openpgp.ReadKeyRing(keyringFileBuffer)
	if err != nil {
		return nil, err
	}
	entity = entityList[0]

	// Get the passphrase and read the private key.
	// Have not touched the encrypted string yet
	passphraseByte := []byte(passphrase)
	log.Println("Decrypting private key using passphrase")
	entity.PrivateKey.Decrypt(passphraseByte)
	for _, subkey := range entity.Subkeys {
		subkey.PrivateKey.Decrypt(passphraseByte)
	}
	log.Println("Finished decrypting private key using passphrase")

	// Decrypt it with the contents of the private key
	md, err := openpgp.ReadMessage(bytes.NewBuffer(encBytes), entityList, nil, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
