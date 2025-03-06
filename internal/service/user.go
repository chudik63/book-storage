package service

import (
	"book-storage/internal/models"
	"book-storage/internal/repository"
	"book-storage/pkg/otp"
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

type EmailService interface {
	SendVerificationEmail(context context.Context, inp *models.SendEmailInput)
}

type UserService struct {
	repo UserRepository
	hash Hasher

	emailService EmailService
}

func NewUserService(repo UserRepository, hash Hasher, es EmailService) *UserService {
	return &UserService{
		repo:         repo,
		hash:         hash,
		emailService: es,
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

	code := otp.GenerateCode()

	user = &models.User{
		Login:            inp.Login,
		Name:             inp.Name,
		Email:            inp.Email,
		Password:         hashedPassword,
		Salt:             string(salt),
		VerificationCode: code,
		CreatedAt:        time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}

	go s.emailService.SendVerificationEmail(ctx, &models.SendEmailInput{
		UserName: inp.Name,
		Email:    inp.Email,
		Code:     code,
	})

	return nil
}
