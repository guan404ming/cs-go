#!/bin/bash

# Build the application
./scripts/build.sh

# Clear the database
rm -f storage/db.json

# Create test input file
cat > test_input.txt << EOF
REGISTER user1
CREATE_LISTING user1 'Phone model 8' 'Black color, brand new' 1000 'Electronics'
GET_LISTING user1 100001
CREATE_LISTING user1 'Black shoes' 'Training shoes' 100 'Sports'
REGISTER user2
REGISTER user2
CREATE_LISTING user2 'T-shirt' 'White color' 20 'Sports'
GET_LISTING user1 100003
GET_CATEGORY user1 'Fashion'
GET_CATEGORY user1 'Sports'
GET_TOP_CATEGORY user1
DELETE_LISTING user1 100003
DELETE_LISTING user2 100003
GET_TOP_CATEGORY user2
DELETE_LISTING user1 100002
GET_TOP_CATEGORY user1
GET_TOP_CATEGORY user3
EOF

# Run the test
cat test_input.txt | ./cs-go.out

# Clean up - remove the test input file
rm -f test_input.txt