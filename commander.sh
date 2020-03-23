#!/bin/sh
rm -rf go-persistence
go build go-persistence.go
sudo cp ./go-persistence /usr/local/bin
sudo groupadd go-persistence
sudo gpasswd -a $USER go-persistence
exec su -l $USER
