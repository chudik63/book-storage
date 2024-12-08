package http

import (
	"book-storage/internal/transport/http/handlers"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func NewBookStorageMux(ctx context.Context) http.Handler {
	router := mux.NewRouter()

	handlers.NewBookHandler(ctx, router)
	handlers.NewUserHandler(ctx, router)

	return router
}
