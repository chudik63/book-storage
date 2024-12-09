package middleware

import (
	"book-storage/pkg/logger"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), logger.RequestID, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
