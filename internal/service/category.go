package service

import (
	"sort"
	"strings"

	"github.com/guan404ming/cs-go/internal/models"
	"github.com/guan404ming/cs-go/internal/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
	listingRepo  *repository.ListingRepository
	userRepo     *repository.UserRepository
}

func NewCategoryService(
	categoryRepo *repository.CategoryRepository,
	listingRepo *repository.ListingRepository,
	userRepo *repository.UserRepository,
) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
		listingRepo:  listingRepo,
		userRepo:     userRepo,
	}
}

func (s *CategoryService) GetCategory(username, category string) ([]models.Listing, error) {
	// 驗證用戶是否存在
	_, err := s.userRepo.GetUser(username)
	if err != nil {
		return nil, err
	}

	// 獲取該類別下的所有商品ID
	listingIDs, err := s.categoryRepo.GetCategoryListings(category)
	if err != nil {
		return nil, err
	}

	// 獲取每個商品的詳細信息
	listings := []models.Listing{}
	for _, id := range listingIDs {
		listing, err := s.listingRepo.GetListing(id)
		if err == nil { // 忽略不存在的商品
			listings = append(listings, listing)
		}
	}

	return listings, nil
}

func (s *CategoryService) GetTopCategory(username string) (string, error) {
	// 驗證用戶是否存在
	_, err := s.userRepo.GetUser(username)
	if err != nil {
		return "", err
	}

	// 獲取所有類別及其商品
	categories, err := s.categoryRepo.GetAllCategories()
	if err != nil || len(categories) == 0 {
		return "Error - no categories found", err
	}

	// Find the maximum count
	maxCount := 0
	for _, listings := range categories {
		count := len(listings)
		if count > maxCount {
			maxCount = count
		}
	}

	// Collect all categories with the maximum count
	var topCategories []string
	for name, listings := range categories {
		if len(listings) == maxCount {
			topCategories = append(topCategories, name)
		}
	}

	// Sort categories lexically
	sort.Strings(topCategories)

	// Join the top categories with comma
	return strings.Join(topCategories, "\n"), nil
}
