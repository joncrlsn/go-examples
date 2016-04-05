package main

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"log"
	"os"
)

// create gpg keys with
// $ gpg --gen-key
// ensure you specify correct paths and passphrase

func main() {

	mySecretString := "this is so very secret!"
	prefix, passphrase := "/home/jon/", "xxxxxxxxx"
	secretKeyring := prefix + ".gnupg/secring.gpg"
	publicKeyring := prefix + ".gnupg/pubring.gpg"

	if 1 == 1 {
		bytes, err := ioutil.ReadFile(prefix + ".password-store/jon/facebook:jon_carlson@writeme.com.gpg")
		if err != nil {
			log.Fatal(err)
		}
		decBytes, err := DecryptBytes(bytes, secretKeyring, passphrase)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("decrypted stuff", string(decBytes))
	}

	// Encrypt string
	encBytes, err := EncryptBytes([]byte(mySecretString), publicKeyring)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt string
	decBytes, err := DecryptBytes(encBytes, secretKeyring, passphrase)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Decrypted Secret:", string(decBytes))
}

func EncryptBytes(origBytes []byte, pubKeyringFile string) ([]byte, error) {

	// Read in public key
	keyringFileBuffer, _ := os.Open(pubKeyringFile)
	defer keyringFileBuffer.Close()
	entityList, err := openpgp.ReadKeyRing(keyringFileBuffer)
	if err != nil {
		return nil, err
	}

	// encrypt string
	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entityList, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(origBytes)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	encryptedBytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return nil, err
	}

	return encryptedBytes, nil
}

func DecryptBytes(encBytes []byte, privKeyringFile string, passphrase string) ([]byte, error) {

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
