package service

import (
	"book-storage/internal/models"
	"book-storage/internal/repository"
	"context"
	"fmt"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	GetByCredentials(ctx context.Context, equatoins repository.Creds) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID int64) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Create(ctx context.Context, user *models.User) (int64, error) {
	if user, _ := s.repo.GetByCredentials(ctx, repository.Creds{"login": user.Login}); user != nil {
		return user.ID, fmt.Errorf("user already exists")
	}

	id, err := s.repo.Create(ctx, user)
	return id, err
}

func (s *UserService) Read(ctx context.Context, userID int64) (*models.User, error) {
	return s.repo.GetByCredentials(ctx, repository.Creds{"id": userID})
}

func (s *UserService) Update(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, userID int64) error {
	return s.repo.Delete(ctx, userID)
}
