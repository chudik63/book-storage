package http

import (
	"book-storage/internal/models"
	"book-storage/pkg/logger"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Response map[string]interface{}

type UserService interface {
	Create(ctx context.Context, user *models.User) (int64, error)
	Read(ctx context.Context, userID int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID int64) error
}

type UserHandler struct {
	userService UserService
	logger      logger.Logger
}

func NewUserHandler(ctx context.Context, mux *mux.Router, service UserService) {
	userHandler := &UserHandler{
		userService: service,
		logger:      logger.GetLoggerFromCtx(ctx),
	}

	mux.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	mux.HandleFunc("/api/users/{id}", userHandler.ReadUser).Methods("GET")
	mux.HandleFunc("/api/users/{id}", userHandler.UpdateUser).Methods("PATCH")
	mux.HandleFunc("/api/users/{id}", userHandler.DeleteUser).Methods("DELETE")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error(r.Context(), "CreateUser: can`t decode request", zap.String("err", err.Error()))
		return
	}

	id, err := h.userService.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error(r.Context(), "CreateUser: can`t create user", zap.String("err", err.Error()))
		return
	}

	response := Response{
		"id": id,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) ReadUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
