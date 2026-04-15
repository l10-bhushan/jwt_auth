package service

import (
	"context"

	"github.com/l10-bhushan/jwt_auth/internal/repository"
)

type UserService struct {
	repo *repository.InMemoryUserRepo
}

func NewUserService(repo *repository.InMemoryUserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

// TODO: Implement Login functionality
func (service *UserService) Login(ctx context.Context, username, password string) error {
	return nil
}

// TODO: Implement Signup functionality
func (service *UserService) SignUp(ctx context.Context, email, username, password string) error {
	return nil
}
