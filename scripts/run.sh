#!/bin/bash

# Make sure to build the latest version
./scripts/build.sh

# Clear the database
rm -f storage/db.json
cat > storage/db.json << EOF
{}
EOF

# Run the application
./cs-go.out "$@" 