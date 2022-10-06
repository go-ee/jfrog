#!/bin/bash

export RELEASES=REALEASES
mkdir $RELEASES

GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static'" -o "$RELEASES/linux/"
GOOS=darwin GOARCH=arm64 go build -ldflags "-extldflags '-static'" -o "$RELEASES/macos/"
GOOS=windows GOARCH=amd64 go build -ldflags "-extldflags '-static'" -o "$RELEASES/windows/"