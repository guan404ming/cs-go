#!/bin/bash

# Make sure to build the latest version
./scripts/build.sh

# Clear the database
rm -f storage/db.json
cat > storage/db.json << EOF
{
    "categories": {},
    "listings": {},
    "users": {}
}
EOF

# Run the application
echo "================================================"
echo "CS-GO CLI is running..."
echo "================================================"
./cs-go.out "$@" 