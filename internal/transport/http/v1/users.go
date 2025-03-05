package v1

import (
	"book-storage/internal/models"
	"book-storage/pkg/email"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) InitUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.POST("/sign-in", h.userSignIn)
		users.POST("/auth/refresh", h.userRefresh)
	}
}

func (h *Handler) userSignUp(c *gin.Context) {
	var inp models.SignUpInput
	if err := c.BindJSON(&inp); err != nil {
		h.errorResponse(c, http.StatusBadGateway, "invalid input")

		return
	}

	if len(inp.Password) > 12 {
		h.errorResponse(c, http.StatusBadGateway, models.ErrPasswordIsTooLong.Error())

		return
	}

	if len(inp.Login) > 12 {
		h.errorResponse(c, http.StatusBadGateway, models.ErrLoginIsTooLong.Error())

		return
	}

	if len(inp.Name) > 20 {
		h.errorResponse(c, http.StatusBadGateway, models.ErrNameIsTooLong.Error())

		return
	}

	if !email.IsValid(inp.Email) {
		h.errorResponse(c, http.StatusBadGateway, models.ErrEmailFormat.Error())

		return
	}

	err := h.userService.SignUp(c.Request.Context(), &inp)

	if err != nil {

		if errors.Is(err, models.ErrLoginAlreadyExists) || errors.Is(err, models.ErrEmailAlreadyRegistered) {
			h.errorResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		h.errorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	h.logs.Info(c.Request.Context(), "User created", zap.String("name", inp.Name), zap.String("login", inp.Login), zap.String("email", inp.Email))

	c.Status(http.StatusCreated)
}

func (h *Handler) userSignIn(c *gin.Context) {

}

func (h *Handler) userRefresh(c *gin.Context) {

}
