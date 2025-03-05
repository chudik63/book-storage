package repository

import (
	"book-storage/internal/database/postgres"
	"book-storage/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

type Creds map[string]interface{}

type UserRepository struct {
	db postgres.DB
}

func NewUserRepository(db postgres.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := sq.Insert("users").
		Columns("login", "name", "email", "password_hash", "password_salt", "created_at").
		Values(user.Login, user.Name, user.Email, user.Password, user.Salt, user.CreatedAt).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		Exec()

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByCredentials(ctx context.Context, equations Creds) (*models.User, error) {
	var user models.User

	sb := sq.Select("id, login, name, password_hash, email").
		From("users")

	for field, value := range equations {
		sb = sb.Where(sq.Eq{field: value})
	}

	row := sb.PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		QueryRow()

	err := row.Scan(&user.ID, &user.Login, &user.Name, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get by credential: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	_, err := sq.Update("public.user").
		Set("name", user.Name).
		Set("password", user.Password).
		Where(sq.Eq{"id": user.ID}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		Exec()

	return err
}

func (r *UserRepository) Delete(ctx context.Context, userID int64) error {
	_, err := sq.Delete("public.users").
		Where(sq.Eq{"id": strconv.FormatInt(userID, 10)}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		Exec()

	return err
}
