package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/guan404ming/cs-go/internal/repository"
	"github.com/guan404ming/cs-go/internal/service"
)

var (
	dbPath          = "storage/db.json"
	userRepo        *repository.UserRepository
	listingRepo     *repository.ListingRepository
	categoryRepo    *repository.CategoryRepository
	userService     *service.UserService
	listingService  *service.ListingService
	categoryService *service.CategoryService
)

func init() {
	// Ensure storage directory exists
	os.MkdirAll(filepath.Dir(dbPath), 0755)

	// Initialize repositories
	userRepo = repository.NewUserRepository(dbPath)
	listingRepo = repository.NewListingRepository(dbPath)
	categoryRepo = repository.NewCategoryRepository(dbPath)

	// Initialize services
	userService = service.NewUserService(userRepo)
	listingService = service.NewListingService(listingRepo, userRepo, categoryRepo)
	categoryService = service.NewCategoryService(categoryRepo, listingRepo, userRepo)
}

func Execute() error {
	if len(os.Args) < 2 {
		return errors.New("No command provided")
	}

	cmd := strings.ToUpper(os.Args[1])
	args := os.Args[2:]

	switch cmd {
	case "REGISTER":
		return handleRegister(args)
	case "CREATE_LISTING":
		return handleCreateListing(args)
	case "DELETE_LISTING":
		return handleDeleteListing(args)
	case "GET_LISTING":
		return handleGetListing(args)
	case "GET_CATEGORY":
		return handleGetCategory(args)
	case "GET_TOP_CATEGORY":
		return handleGetTopCategory(args)
	default:
		return fmt.Errorf("Unknown command: %s", cmd)
	}
}
