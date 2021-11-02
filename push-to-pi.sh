#!/bin/sh

env GOOS=linux GOARCH=arm GOARM=7 go build
scp hcp pi:/home/pi/.local/bin
rm hcp
