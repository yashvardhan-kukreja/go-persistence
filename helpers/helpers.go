package helpers

import (
	"errors"
	"github.com/boltdb/bolt"
)

//CreateBucket (db *bolt.DB, name string) - Helper for creating a bucket
func CreateBucket(db *bolt.DB, name string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, createError := tx.CreateBucket([]byte(name))
		if createError != nil {
			return createError
		} 
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

//DeleteBucket (db *bolt.DB, name string) - Helper for deleting a bucket
func DeleteBucket(db *bolt.DB, name string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		delErr := tx.DeleteBucket([]byte(name))
		if delErr != nil {
			return delErr
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

//AddKey (db *bolt.DB, bucketName string, key string, value string) - Helper for adding a key-value pair to a bucket
func AddKey(db *bolt.DB, bucketName string, key string, value string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		keyErr := tx.Bucket([]byte(bucketName)).Put([]byte(key), []byte(value))
		if keyErr != nil {
			return keyErr
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

//RemoveKey (db *bolt.DB, bucketName string, key string) - Helper for removing a key value pair from a bucket
func RemoveKey(db *bolt.DB, bucketName string, key string) error {
	err := db.Update(func (tx *bolt.Tx) error {
		cursor := tx.Bucket([]byte(bucketName)).Cursor()
		k, _ := cursor.Seek([]byte(key))

		if k == nil || string(k) != key {
			return errors.New("Key not found in the database")
		}
		delErr := tx.Bucket([]byte(bucketName)).Delete([]byte(key))
		if delErr != nil {
			return delErr
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

//GetValue (db *bolt.DB, bucketName string, key string) - Helper for getting the value of a key from a bucket
func GetValue(db *bolt.DB, bucketName string, key string) (*string, error) {

	outputValue := ""
	err := db.View(func(tx *bolt.Tx) error {
		cursor := tx.Bucket([]byte(bucketName)).Cursor()
		k, v := cursor.Seek([]byte(key))

		if k == nil || string(k) != key {
			return errors.New("Key not found in the database")
		}

		outputValue = string(v)
		return nil

		//outputValue = string(tx.Bucket([]byte(bucketName)).Get([]byte(key)))
		//return nil
	})

	if err != nil {
		return nil, err
	}
	return &outputValue, nil
}



