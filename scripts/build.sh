#!/bin/bash

echo "Building cs-go CLI..."
go build -o cs-go main.go

if [ $? -eq 0 ]; then
    echo "Build successful!"
else
    echo "Build failed!"
    exit 1
fi 