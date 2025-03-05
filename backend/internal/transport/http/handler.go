package http

import (
	"book-storage/internal/config"
	v1 "book-storage/internal/transport/http/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService v1.UserService
}

func NewHandler(us v1.UserService) *Handler {
	return &Handler{
		userService: us,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.userService)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
