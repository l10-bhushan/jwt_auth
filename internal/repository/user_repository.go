package repository

import (
	"context"
	"time"

	"github.com/l10-bhushan/jwt_auth/internal/dto"
	"github.com/l10-bhushan/jwt_auth/internal/model"
)

// User repository interface that keeps method definition
type UserRepository interface {
	Login(ctx context.Context, username string, password string) (dto.UserCreationSuccess, error)
	SignUp(ctx context.Context, username string, password string, email string) (dto.UserCreationSuccess, error)
}

// Using in-memory user repo we will replace it later
// TODO: Add database later
type InMemoryUserRepo struct {
	data map[string]model.User
}

// Constructor function to create instance of in-memory repo
func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data: make(map[string]model.User),
	}
}

// Login method implementation
func (repo *InMemoryUserRepo) Login(ctx context.Context, username string, password string) error {
	return nil
}

// Signup method implementation
func (repo *InMemoryUserRepo) SignUp(ctx context.Context, id string, email string, username string, password string) error {
	repo.data[id] = model.User{
		Id:        id,
		Email:     email,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	return nil
}
