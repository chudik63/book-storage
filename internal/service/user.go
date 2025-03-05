package service

import (
	"book-storage/internal/models"
	"book-storage/internal/repository"
	"context"
	"errors"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByCredentials(ctx context.Context, equations repository.Creds) (*models.User, error)
}

type Hasher interface {
	GenerateSalt() (string, error)
	HashPassword(password, salt string) string
	DoPasswordsMatch(hashedPassword, currPassword, salt string) bool
}

type UserService struct {
	repo UserRepository
	hash Hasher
}

func NewUserService(repo UserRepository, hash Hasher) *UserService {
	return &UserService{
		repo: repo,
		hash: hash,
	}
}

func (s *UserService) SignUp(ctx context.Context, inp *models.SignUpInput) error {
	user, err := s.repo.GetByCredentials(ctx, repository.Creds{"login": inp.Login})
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return err
	}
	if user != nil {
		return models.ErrLoginAlreadyExists
	}

	user, err = s.repo.GetByCredentials(ctx, repository.Creds{"email": inp.Email})
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return err
	}
	if user != nil {
		return models.ErrEmailAlreadyRegistered
	}

	salt, err := s.hash.GenerateSalt()
	if err != nil {
		return err
	}

	hashedPassword := s.hash.HashPassword(inp.Password, salt)

	user = &models.User{
		Login:     inp.Login,
		Name:      inp.Name,
		Email:     inp.Email,
		Password:  hashedPassword,
		Salt:      string(salt),
		CreatedAt: time.Now(),
	}

	return s.repo.Create(ctx, user)
}
