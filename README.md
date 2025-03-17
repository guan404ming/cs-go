# 🛒 CS-GO: Cloud Shop CLI Application

CS-GO is a command-line interface (CLI) application that allows users to register, create product listings, delete products, query products, and categories.

## ✨ Features

- 👤 User registration (REGISTER)
- 📦 Create product listings (CREATE_LISTING)
- 🗑️ Delete product listings (DELETE_LISTING)
- 🔍 Query product details (GET_LISTING)
- 📋 Query products by category (GET_CATEGORY)
- 📊 Query the most popular category (GET_TOP_CATEGORY)

## 📋 Requirements

- Go 1.16 or higher

## 📝 Implementation Details

- 🔢 Product IDs are assigned sequentially starting from 100001
- ⏰ Current system time is used for timestamps when displaying products

## 🚀 Build and Run

### 🔨 Build

```bash
./scripts/build.sh
```

### 🏃‍♂️ Run

```bash
./scripts/run.sh
```

Or run directly:

```bash
./cs-go
```

## 📖 Usage

### 👤 Register a user

```
REGISTER <username>
```

### 📦 Create a product listing

```
CREATE_LISTING <username> <title> <description> <price> <category>
```

### 🗑️ Delete a product listing

```
DELETE_LISTING <username> <listing_id>
```

### 🔍 Query product details

```
GET_LISTING <username> <listing_id>
```

### 📋 Query products by category

```
GET_CATEGORY <username> <category>
```

### 📊 Query the most popular category

```
GET_TOP_CATEGORY <username>
```

## 🏗️ Architecture

This project uses a layered architecture to ensure separation of concerns, making it easy to maintain and extend.

### 🧱 Layer Design

- **cmd**: CLI command processing
- **internal/service**: Service layer, business logic processing
- **internal/repository**: Data access layer
- **internal/models**: Structure definitions
- **pkg**: Shared libraries
- **storage**: Persistent storage

## 🧪 Testing

Run the test script:

```bash
./scripts/test.sh
``` 