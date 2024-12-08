package handlers

import (
	"book-storage/pkg/logger"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	// bookService *services.BookService
	l logger.Logger
}

func NewBookHandler(ctx context.Context, mux *mux.Router) {
	bookHandler := &BookHandler{
		l: logger.GetLoggerFromCtx(ctx),
	}

	mux.HandleFunc("/api/books", bookHandler.CreateBook).Methods("POST")
	mux.HandleFunc("/api/books/{id}", bookHandler.ReadBook).Methods("GET")
	mux.HandleFunc("/api/books/{id}", bookHandler.UpdateBook).Methods("PATCH")
	mux.HandleFunc("/api/books/{id}", bookHandler.DeleteBook).Methods("DELETE")
	mux.HandleFunc("/api/books", bookHandler.ListBook).Methods("GET")
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) ReadBook(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) ListBook(w http.ResponseWriter, r *http.Request) {

}
