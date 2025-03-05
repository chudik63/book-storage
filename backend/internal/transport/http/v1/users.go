package v1

import (
	"book-storage/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var inp signUpInput
	if err := c.BindJSON(&inp); err != nil {
		h.errorResponse(c, http.StatusBadGateway, "invalid input")

		return
	}

	err := h.userService.SignUp(c.Request.Context(), &models.User{
		Login:    inp.Login,
		Name:     inp.Name,
		Password: inp.Password,
		Email:    inp.Email,
	})

	if err != nil {

		if errors.Is(err, models.ErrLoginAlreadyExists) || errors.Is(err, models.ErrEmailAlreadyRegistered) {
			h.errorResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		h.errorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) userSignIn(c *gin.Context) {

}

func (h *Handler) userRefresh(c *gin.Context) {

}
