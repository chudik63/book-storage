package http

import (
	"book-storage/internal/config"
	v1 "book-storage/internal/transport/http/v1"
	"book-storage/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService v1.UserService

	logs logger.Logger
}

func NewHandler(us v1.UserService, l logger.Logger) *Handler {
	return &Handler{
		userService: us,
		logs:        l,
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
	handlerV1 := v1.NewHandler(h.userService, h.logs)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
