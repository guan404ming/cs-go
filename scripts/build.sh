#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

echo "Building cs-go CLI..."
go build -o cs-go.out main.go

if [ $? -eq 0 ]; then
    echo "Build successful!"
    chmod +x cs-go.out
else
    echo "Build failed!"
    exit 1
fi 