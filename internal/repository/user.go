package repository

import (
	"book-storage/internal/models"
	"book-storage/pkg/postgres"
	"context"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db postgres.DB
}

func NewUserRepository(db postgres.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (int64, error) {
	var id int64

	err := sq.Insert("public.users").
		Columns("name", "password").
		Values(user.Name, user.Password).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		QueryRow().
		Scan(&id)

	return id, err
}

func (r *UserRepository) Read(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User

	err := sq.Select("*").
		From("public.users").
		Where(sq.Eq{"id": strconv.FormatInt(userID, 10)}).
		PlaceholderFormat(sq.Dollar).
		RunWith(r.db).
		QueryRow().
		Scan(&user.ID, &user.Name, &user.Password)

	return &user, err
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
