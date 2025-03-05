package v1

import (
	"book-storage/internal/models"
	"context"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	SignUp(ctx context.Context, user *models.SignUpInput) error
	Read(ctx context.Context, userID int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID int64) error
}

type Handler struct {
	userService UserService
}

func NewHandler(us UserService) *Handler {
	return &Handler{
		userService: us,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitUserRoutes(v1)
	}
}

func (h *Handler) errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, models.Response{Message: msg})
}
