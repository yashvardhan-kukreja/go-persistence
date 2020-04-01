package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"github.com/boltdb/bolt"
	"github.com/yashvardhan-kukreja/go-persistence/helpers"
)

func main() {
	createBucket := flag.Bool("create-bucket", false, "Operation for creating a bucket")
	deleteBucket := flag.Bool("delete-bucket", false, "Operation for removing/deleting a bucket with all its contents")
	addKey := flag.Bool("add-key", false, "Operation for adding a key to a bucket")
	removeKey := flag.Bool("remove-key", false, "Operation for removing a key from a bucket")
	getValue := flag.Bool("get-value", false, "Operation for fetching a key-value from a bucket by the provided keyname")
	randomItem := flag.Bool("random-item", false, "Operation for fetching a key-value pair from the provided bucket name")

	bucketName := flag.String("bucket-name", "", "The name of bucket you want to apply the operation on")
	keyName := flag.String("key", "", "The key for which you are performing the operation")
	valueName := flag.String("value", "", "The key for which you are performing the operation")

	flag.Parse()

	db, dbOpenErr := bolt.Open("./go-persistence.db", 0640, nil)
	if dbOpenErr != nil {
		log.Fatal(dbOpenErr)
		return
	}
	defer db.Close()

	if *createBucket {
		if *bucketName == "" {
			log.Fatal(errors.New("bucket name not provided for creating the bucket"))
			return
		}
		err := helpers.CreateBucket(db, *bucketName)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Created a bucket with the name: %s\n", *bucketName)
		return
	}

	if *deleteBucket {
		if *bucketName == "" {
			log.Fatal(errors.New("bucket name not provided for deleting the bucket"))
			return
		}
		err := helpers.DeleteBucket(db, *bucketName)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Deleted the bucket with the name: %s\n", *bucketName)
		return
	}

	if *addKey {
		if *bucketName == "" || *keyName == "" || *valueName == "" {
			log.Fatal(errors.New("bucket name or key name or value not provided for adding the key-value pair"))
			return
		}
		err := helpers.AddKey(db, *bucketName, *keyName, *valueName)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("For the bucket: %s\nAdded key: %s  AND value: %s\n", *bucketName, *keyName, *valueName)
		return
	}

	if *removeKey {
		if *bucketName == "" || *keyName == "" {
			log.Fatal(errors.New("bucket name or key name not provided for removing the key"))
			return
		}
		err := helpers.RemoveKey(db, *bucketName, *keyName)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("From the bucket: %s\nRemoved the key: %s\n", *bucketName, *keyName)
		return
	}

	if *getValue {
		if *bucketName == "" || *keyName == "" {
			log.Fatal(errors.New("bucket name or key name not provided for removing the key"))
			return
		}
		outputValuePtr, err := helpers.GetValue(db, *bucketName, *keyName)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Bucket: %s\nKey: %s\nValue: %s\n", *bucketName, *keyName, *outputValuePtr)
		return
	}

	if *randomItem {
		if *bucketName == "" {
			log.Fatal(errors.New("bucket name not provided for fetching the random key-value pair from"))
			return
		}
		randomKeyPtr, randomValuePtr, err := helpers.RandomItem(db, *bucketName)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Bucket: %s\nAnd your random item from the above bucket is-----\nKey: %s\nValue: %s\n", *bucketName, *randomKeyPtr, *randomValuePtr)
		return
	}
}
