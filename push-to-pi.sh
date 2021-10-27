#!/bin/sh

env GOOS=linux GOARCH=arm GOARM=7 go build
scp hcp pi:/home/pi
rm hcp
