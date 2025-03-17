package service

import (
	"github.com/guan404ming/cs-go/internal/models"
	"github.com/guan404ming/cs-go/internal/repository"
)

type ListingService struct {
	listingRepo  *repository.ListingRepository
	userRepo     *repository.UserRepository
	categoryRepo *repository.CategoryRepository
}

func NewListingService(
	listingRepo *repository.ListingRepository,
	userRepo *repository.UserRepository,
	categoryRepo *repository.CategoryRepository,
) *ListingService {
	return &ListingService{
		listingRepo:  listingRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *ListingService) CreateListing(username, title, description string, price float64, category string) (string, error) {
	// 驗證用戶是否存在
	_, err := s.userRepo.GetUser(username)
	if err != nil {
		return "", err
	}

	// 創建商品
	listingID, err := s.listingRepo.CreateListing(title, description, price, category, username)
	if err != nil {
		return "", err
	}

	// 將商品添加到用戶的商品列表中
	if err := s.userRepo.AddListingToUser(username, listingID); err != nil {
		return "", err
	}

	// 將商品添加到類別中
	if err := s.categoryRepo.AddListingToCategory(category, listingID); err != nil {
		return "", err
	}

	return listingID, nil
}

func (s *ListingService) DeleteListing(username, listingID string) error {
	// 驗證用戶是否存在
	_, err := s.userRepo.GetUser(username)
	if err != nil {
		return err
	}

	// 獲取商品信息
	listing, err := s.listingRepo.GetListing(listingID)
	if err != nil {
		return err
	}

	// 刪除商品
	if err := s.listingRepo.DeleteListing(listingID, username); err != nil {
		return err
	}

	// 從用戶的商品列表中移除
	if err := s.userRepo.RemoveListingFromUser(username, listingID); err != nil {
		return err
	}

	// 從類別中移除
	if err := s.categoryRepo.RemoveListingFromCategory(listing.Category, listingID); err != nil {
		return err
	}

	return nil
}

func (s *ListingService) GetListing(username, listingID string) (models.Listing, error) {
	// 驗證用戶是否存在
	_, err := s.userRepo.GetUser(username)
	if err != nil {
		return models.Listing{}, err
	}

	// 獲取商品信息
	return s.listingRepo.GetListing(listingID)
}
