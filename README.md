# go-persistence

## Introduction: 
A friendly CLI to provide you a bucket-based persistent key-value store<br>
It stores key-value pairs in dedicated buckets.<br> So, you can have two key-value pairs with same key name but in different buckets.<br>
For example, there can be a bucket named "Alex" where there is a key-value pair -> (**"name"**, "Alex")<br>
And there can be a bucket named "Bob" where there is a key-value pair -> (**"name"**, "Bob")

------

## Pre-requisites:
Please have "go" installed in your computer. That's it!

------

## How to set it up:
 - Download the github repo with the following command:<br>
 ```git clone https://www.github.com/yashvardhan-kukreja/go-persistence```
 - Then, run the following command:<br>
 ```bash commander.sh```

 - And, that's it !

-----

## How to run it:
 
 The command expects one of the following operations:
 - **create-bucket** -> Operation for creating a bucket
 - **delete-bucket** -> Operation for removing/deleting a bucket with all its contents
 - **add-key** -> Operation for adding a key-value pair to a bucket
 - **remove-key**  -> Operation for removing a key from a bucket
 - **get-value** -> Operation for fetching a key-value from a bucket by the provided keyname
 <br>

The command expects the following fields on the basis of the respective operation:
 - **bucket-name** -> The name of the bucket you are performing operation on
 - **key** -> Name of the key (in the bucket) you are performing operation on
 - **value** -> Value of a key for which you want add a key-value pair


**Some Examples command:**<br>

Creating a bucket
 ```
 ► go-persistence --create-bucket --bucket-name b
Created a bucket with the name: b
 ```

 Adding a key-value pair from a bucket
 ```
 ► go-persistence --add-key --bucket-name b --key name --value yash
For the bucket: b
Added key: name  AND value: yash
 ```

 Removing a key-value pair from a bucket
 ```
 ► go-persistence --remove-key --bucket-name b --key name --value yash1
From the bucket: b
Removed the key: name
 ```

 Trying to remove a key-value pair for which the key does not exist
 ```
 ► go-persistence --remove-key --bucket-name b --key name --value yash1
2020/03/24 03:54:30 Key not found in the database
 ```

 Deleting a bucket
 ```
 ► go-persistence --delete-bucket --bucket-name b
Deleted the bucket with the name: b
 ```

-----