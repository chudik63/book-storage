package service

import (
	"book-storage/internal/models"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	Read(ctx context.Context, userID int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID int64) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo}
}
