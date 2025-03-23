# 🛒 CS-GO: Cloud Shop CLI Application

CS-GO is a command-line interface (CLI) application that allows users to register, create product listings, delete products, query products, and categories.

## ✨ Features

- 👤 User registration (REGISTER)
- 📦 Create product listings (CREATE_LISTING)
- 🗑️ Delete product listings (DELETE_LISTING)
- 🔍 Query product details (GET_LISTING)
- 📋 Query products by category (GET_CATEGORY)
- 📊 Query the most popular category (GET_TOP_CATEGORY)

## 🚀 Development (Requirements: Go 1.24 or higher)

use `./scripts/*` to build the application

```bash
# Build the application
./scripts/build.sh

# Run the application
./scripts/run.sh

# Run the tests
./scripts/test.sh
```

or use `make` to build, run, and test the application

```bash
# Build the application
make build

# Run the application
make run

# Run the tests
make test
```

## 📖 Usage

- Register a user - `REGISTER <username>`

- Create a product listing - `CREATE_LISTING <username> <title> <description> <price> <category>`

- Delete a product listing - `DELETE_LISTING <username> <listing_id>`

- Query product details - `GET_LISTING <username> <listing_id>`

- Query products by category - `GET_CATEGORY <username> <category>`

- Query the most popular category - `GET_TOP_CATEGORY <username>`

## 🏗️ Architecture

This project uses a layered architecture to ensure separation of concerns, making it easy to maintain and extend.

- **cmd**: CLI command processing
- **internal/service**: Service layer, business logic processing
- **internal/repository**: Data access layer
- **internal/models**: Structure definitions
- **storage**: Persistent storage
