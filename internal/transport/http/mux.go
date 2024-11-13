package http

import (
	"book-storage/internal/transport/http/handlers"
	"book-storage/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

func NewBookStorageMux(logs logger.Logger) http.Handler {
	router := mux.NewRouter()

	handlers.NewBookHandler(logs, router)

	return router
}
