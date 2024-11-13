package handlers

import (
	"book-storage/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	// bookService *services.BookService
	l logger.Logger
}

func NewBookHandler(logs logger.Logger, mux *mux.Router) {
	bookHandler := &BookHandler{
		l: logs,
	}

	mux.HandleFunc("/api/books", bookHandler.CreateBook).Methods("POST")
	mux.HandleFunc("/api/books/{id}", bookHandler.ReadBook).Methods("GET")
	mux.HandleFunc("/api/books/{id}", bookHandler.UpdateBook).Methods("PATCH")
	mux.HandleFunc("/api/books/{id}", bookHandler.DeleteBook).Methods("DELETE")
	mux.HandleFunc("/api/books", bookHandler.ListBook).Methods("GET")
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	h.l.Info("create book handled")
}

func (h *BookHandler) ReadBook(w http.ResponseWriter, r *http.Request) {
	h.l.Info("read book handled")
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	h.l.Info("update book handled")
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	h.l.Info("delete book handled")
}

func (h *BookHandler) ListBook(w http.ResponseWriter, r *http.Request) {
	h.l.Info("list book handled")
}
