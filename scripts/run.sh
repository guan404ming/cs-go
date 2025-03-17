#!/bin/bash

# Make sure to build the latest version
./scripts/build.sh

# Run the application
./cs-go "$@" 