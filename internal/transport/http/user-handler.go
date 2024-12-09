package http

import (
	"book-storage/pkg/logger"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type UserService interface {
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
