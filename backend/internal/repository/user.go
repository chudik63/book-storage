package repository

import (
	"book-storage/internal/database/postgres"
	"book-storage/internal/models"
	"context"
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
		Columns("login", "name", "password", "email").
		Values(user.Name, user.Login, user.Password, user.Email).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		Exec()

	return err
}

func (r *UserRepository) GetByCredentials(ctx context.Context, equations Creds) (*models.User, error) {
	var user models.User

	rows, err := sq.Select("id, login, name, password, email").
		From("users").
		Where(equations).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		Query()

	if err != nil {
		return nil, err
	}

	rows.Scan(&user.ID, &user.Login, &user.Name, &user.Password, &user.Email)

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
