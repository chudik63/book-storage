package handlers

import (
	"book-storage/pkg/logger"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	// bookService *services.BookService
	l logger.Logger
}

func NewUserHandler(ctx context.Context, mux *mux.Router) {
	bookHandler := &BookHandler{
		l: logger.GetLoggerFromCtx(ctx),
	}

	mux.HandleFunc("/api/users", bookHandler.CreateBook).Methods("POST")
	mux.HandleFunc("/api/users/{id}", bookHandler.ReadBook).Methods("GET")
	mux.HandleFunc("/api/users/{id}", bookHandler.UpdateBook).Methods("PATCH")
	mux.HandleFunc("/api/users/{id}", bookHandler.DeleteBook).Methods("DELETE")
}

func (h *BookHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) ReadUser(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
