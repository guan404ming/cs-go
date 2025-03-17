package service

import (
	"github.com/guan404ming/cs-go/internal/models"
	"github.com/guan404ming/cs-go/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(username string) error {
	return s.userRepo.CreateUser(username)
}

func (s *UserService) ValidateUser(username string) (models.User, error) {
	return s.userRepo.GetUser(username)
}
