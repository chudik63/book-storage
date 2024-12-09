package http

import (
	"book-storage/internal/models"
	"book-storage/pkg/logger"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	Read(ctx context.Context, userID int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID int64) error
}

type UserHandler struct {
	userService UserService
	l           logger.Logger
}

func NewUserHandler(ctx context.Context, mux *mux.Router, service UserService) {
	userHandler := &UserHandler{
		userService: service,
		l:           logger.GetLoggerFromCtx(ctx),
	}

	mux.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	mux.HandleFunc("/api/users/{id}", userHandler.ReadUser).Methods("GET")
	mux.HandleFunc("/api/users/{id}", userHandler.UpdateUser).Methods("PATCH")
	mux.HandleFunc("/api/users/{id}", userHandler.DeleteUser).Methods("DELETE")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) ReadUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
