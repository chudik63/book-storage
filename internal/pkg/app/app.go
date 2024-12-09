package app

import (
	"book-storage/internal/config"
	"book-storage/internal/repository"
	"book-storage/internal/service"
	"book-storage/internal/transport/http"
	"book-storage/pkg/logger"
	"book-storage/pkg/postgres"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *config.Config) {
	logs := logger.GetLoggerFromCtx(ctx)

	db := postgres.New(ctx, &cfg.Config)

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository)

	router := mux.NewRouter()

	http.NewBookHandler(ctx, router)
	http.NewUserHandler(ctx, router, userService)

	httpServer := http.NewServer(cfg, router)

	go func() {
		if err := httpServer.Run(); err != nil {
			logs.Error(ctx, "failed starting the server", zap.String("err: ", err.Error()))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	if err := httpServer.Stop(); err != nil {
		logs.Error(ctx, "failed shutting down the server", zap.String("err: ", err.Error()))
	}
}
