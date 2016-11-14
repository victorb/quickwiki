#! /usr/bin/env bash

set -e

COMMIT=$(git rev-parse HEAD)
VERSION=${1:-$COMMIT}

LD_FLAGS="-X github.com/victorbjelkholm/quickwiki/cmd.version=${VERSION}"

echo "## Building version '$VERSION'"
echo "## Building Linux version..."
GOOS=linux time go build -ldflags "$LD_FLAGS" -a -installsuffix cgo -o "dist/quickwiki-linux" main.go
echo
echo "## Building OSX version..."
GOOS=darwin time go build -ldflags "$LD_FLAGS" -a -installsuffix cgo -o "dist/quickwiki-darwin" main.go
echo
echo "## Building Windows version"
GOOS=windows time go build -ldflags "$LD_FLAGS" -a -installsuffix cgo -o "dist/quickwiki-windows" main.go
echo

echo "## Adding to IPFS"
HASH=$(ipfs add -rq dist | tail -n1)
echo "## Added $HASH"
echo "Local: http://localhost:8080/ipfs/$HASH"
echo "Gateway: https://ipfs.io/ipfs/$HASH"
