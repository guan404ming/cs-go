# CS-GO Project Report

## Project Overview

CS-GO (Cloud Shop CLI Application) is a command-line interface application for managing an online marketplace. The application enables users to register accounts, create and manage product listings, browse items by category, and perform various marketplace operations through a simple text-based interface.

GitHub Link: <https://github.com/guan404ming/cs-go>

## Execution Environment and Requirementsß

- **Operating System**: macOS/Linux
- **Go Version**: Go 1.24 or higher
- **Storage**: Local JSON file-based storage
- **Build Tools**: Make, shell scripts

## Package Dependencies and Versions

This project uses only standard library packages (no external dependencies):

| Package              | Version          | Purpose                        |
|----------------------|------------------|---------------------------------|
| fmt                  | Go 1.24 std lib  | Formatted I/O                  |
| os                   | Go 1.24 std lib  | Operating system functionality |
| bufio                | Go 1.24 std lib  | Buffered I/O                   |
| strings              | Go 1.24 std lib  | String manipulation            |
| encoding/json        | Go 1.24 std lib  | JSON encoding/decoding         |
| time                 | Go 1.24 std lib  | Time utilities                 |
| errors               | Go 1.24 std lib  | Error handling                 |
| path/filepath        | Go 1.24 std lib  | File path manipulation         |

## Project Structure

```shell
.
├── cmd/                  # Command handlers
│   ├── root.go           # Root command functionality
│   ├── register.go       # User registration
│   ├── create.go         # Product listing creation
│   ├── delete.go         # Product deletion
│   ├── query.go          # Product and category querying
│   └── category.go       # Category management
├── internal/             # Internal packages
│   ├── models/           # Data models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic layer
├── scripts/              # Helper scripts
│   ├── build.sh          # Build script
│   ├── run.sh            # Run script
│   └── test.sh           # Test script
├── storage/              # Data storage
│   └── db.json           # JSON database file
├── main.go               # Application entry point
├── go.mod                # Go module file
├── Makefile              # Project build automation
└── REPORT.md             # This report
```

## Features

1. **User Management**
   - Register new users

2. **Listing Management**
   - Create product listings with title, description, price, and category
   - Delete product listings

3. **Query Functionality**
   - Query specific product details
   - Browse products by category
   - Find the most popular category

4. **Command Interface**
   - Interactive shell-like interface
   - Command-line argument support

## Build and Run Instructions

### Using Makefile

The project includes a standard Makefile with the following targets:

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Clean build artifacts
make clean

# Clean database
make clean-db

# Format code
make fmt

# Run go vet for static analysis
make vet

# Build release version
make release

# Show all available commands
make help
```

### Using Scripts

Alternatively, you can use the provided shell scripts:

```bash
# Build the application
./scripts/build.sh

# Run the application
./scripts/run.sh

# Run tests
./scripts/test.sh
```

## Command Usage

When running the application, you can use the following commands:

1. **Register a user**

   ```shell
   REGISTER <username>
   ```

2. **Create a product listing**

   ```shell
   CREATE_LISTING <username> <title> <description> <price> <category>
   ```

3. **Delete a product listing**

   ```shell
   DELETE_LISTING <username> <listing_id>
   ```

4. **Query product details**

   ```shell
   GET_LISTING <username> <listing_id>
   ```

5. **Query products by category**

   ```shell
   GET_CATEGORY <username> <category>
   ```

6. **Query the most popular category**

   ```shell
   GET_TOP_CATEGORY <username>
   ```

## Architecture

The project follows a layered architecture:

1. **Presentation Layer** (`cmd/` package)
   - Handles user input parsing
   - Formats output for the command line

2. **Service Layer** (`internal/service/` package)
   - Implements business logic
   - Coordinates between repositories

3. **Repository Layer** (`internal/repository/` package)
   - Manages data access
   - Handles JSON persistence

4. **Model Layer** (`internal/models/` package)
   - Defines data structures
   - Implements data validation

## Data Persistence

The application uses a simple JSON file-based storage system:

- Data is stored in `storage/db.json`
- Each model type (users, listings, categories) is stored as a collection
- The repository layer handles all CRUD operations

## Testing

The project includes a test script (`scripts/test.sh`) that builds the application and runs a series of commands to validate functionality. The tests verify:

1. User registration
2. Product listing creation
3. Product retrieval
4. Category queries
5. Error handling

## Future Improvements

Potential areas for enhancement:

1. Add user authentication with passwords
2. Implement search functionality
3. Add listing reviews and ratings
4. Support for listing images
5. Database migration to a more robust solution (e.g., SQLite)

## Submission Contents

The submitted archive contains:

- Complete source code
- Git history
- Standard Makefile for building the project
- build.sh script for building
- run.sh script for running the application
- test.sh script for testing

**Notes**:

1. All necessary files to build and run the project are included.
2. The Makefile provides standard targets (build, run, test, etc.)
3. The scripts directory contains shell scripts that can be used as alternatives to the Makefile

---
