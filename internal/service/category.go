package service

import (
	"sort"

	"github.com/wesley/cloudShop/internal/models"
	"github.com/wesley/cloudShop/internal/repository"
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
	if err != nil {
		return "", err
	}

	if len(categories) == 0 {
		return "No categories found", nil
	}

	// 計算每個類別的商品數量
	type categoryCount struct {
		name  string
		count int
	}

	var counts []categoryCount
	for name, listings := range categories {
		counts = append(counts, categoryCount{name, len(listings)})
	}

	// 按商品數量排序
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	// 返回商品數量最多的類別
	return counts[0].name, nil
}
