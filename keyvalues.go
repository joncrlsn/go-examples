package main

//
// Wraps boltdb with methods that assume strings for the keys and values
//

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type KeyValueStore struct {
	db *bolt.DB
}

type KeyValue struct {
	Key   string
	Value string
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

	// Put a string in the bucket
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
			fmt.Printf("Good work. Bucket %s does not contain key %s\n", bucket, "bogus-key")
		} else {
			fmt.Printf("Error getting from bucket: %s\n", err)
		}
	}

	// Put a string in the bucket
	err = db.PutString(bucket, "jon1", "42")
	err = db.PutString(bucket, "jon2", "43")
	err = db.PutString(bucket, "jon3", "44")
	if err != nil {
		fmt.Printf("Error putting to bucket: %s", err)
	}

	kvs, err := db.GetKeyValues(bucket, "jon")
	if err != nil {
		fmt.Printf("Error getting from bucket: %s\n", err)
	}

	fmt.Printf("Pulled %d key values with prefix jon\n", len(kvs))
	for _, kv := range kvs {
		fmt.Printf("Key: %s Value: %s\n", kv.Key, kv.Value)
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

// GetKeyValues returns the keys that match the given prefix (with their values)
func (kvs *KeyValueStore) GetKeyValues(bucket string, keyPrefix string) ([]KeyValue, error) {

	var returnValue = []KeyValue{}
	// Partial key match example
	kvs.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		c := tx.Bucket([]byte(bucket)).Cursor()

		prefix := []byte(keyPrefix)
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			kv := KeyValue{string(k), string(v)}
			//fmt.Printf("key=%s, value=%s\n", k, v)
			returnValue = append(returnValue, kv)
		}

		return nil
	})

	return returnValue, nil
}
