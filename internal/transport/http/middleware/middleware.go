package middleware

import (
	"book-storage/pkg/logger"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.Request.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		ctx := context.WithValue(c.Request.Context(), logger.RequestID, reqID)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}

}
