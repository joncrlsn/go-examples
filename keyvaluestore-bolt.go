package main

//
// Wraps boltdb with methods that assume strings for the keys and values
//

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type KeyValueStore struct {
	db *bolt.DB
}

const (
	bucket = "MyBucket"
)

func main() {

	db := NewKeyValueStore("bolt.db.deleteme")
	err := db.CreateBucket(bucket)
	if err != nil {
		fmt.Printf("Error creating bucket: %s", err)
	}

	// Put a string to the bucket
	err = db.PutString(bucket, "answer", "42")
	if err != nil {
		fmt.Printf("Error putting to bucket: %s", err)
	}

	// Get a known value
	var value string
	value, err = db.GetString(bucket, "answer")
	if err != nil {
		fmt.Printf("Error getting from bucket: %s\n", err)
	}
	fmt.Println("The answer is", value)

	// Get a known missing value
	value, err = db.GetString(bucket, "bogus-key")
	if err != nil {
		if err.Error() == "DoesNotExist" {
			fmt.Printf("Bucket %s does not contain key %s\n", bucket, "bogus-key")
		} else {
			fmt.Printf("Error getting from bucket: %s\n", err)
		}
	} else {
		fmt.Println("The answer is", value)
	}
}

func NewKeyValueStore(fileName string) *KeyValueStore {
	boltdb, err := bolt.Open(fileName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return &KeyValueStore{db: boltdb}
}

func (kvs *KeyValueStore) CreateBucket(name string) error {
	return kvs.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}

// PutString is a boltdb wrapper that assumes both the key and value are strings
func (kvs *KeyValueStore) PutString(bucket string, key string, value string) error {
	return kvs.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

// GetString is a boltdb wrapper that assumes both the key and value are strings
func (kvs *KeyValueStore) GetString(bucket string, key string) (string, error) {
	var value string
	err := kvs.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		if v == nil {
			return errors.New("DoesNotExist")
		}
		value = string(v)
		return nil
	})
	if err != nil {
		return "", err
	}
	return value, nil
}
