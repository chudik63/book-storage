package v1

import (
	"book-storage/internal/models"
	"book-storage/internal/transport/http/middleware"
	"book-storage/pkg/logger"
	"context"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	SignUp(ctx context.Context, user *models.SignUpInput) error
}

type Handler struct {
	userService UserService

	logs logger.Logger
}

func NewHandler(us UserService, l logger.Logger) *Handler {
	return &Handler{
		userService: us,
		logs:        l,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitUserRoutes(v1)
	}

	v1.Use(middleware.Middleware())
}

func (h *Handler) errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, models.Response{Message: msg})
}
