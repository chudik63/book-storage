package app

import (
	"book-storage/internal/config"
	"book-storage/internal/database/postgres"
	"book-storage/internal/repository"
	"book-storage/internal/server"
	"book-storage/internal/service"
	transport "book-storage/internal/transport/http"
	"book-storage/pkg/email/smtp"
	"book-storage/pkg/hasher"
	"book-storage/pkg/logger"
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *config.Config) {
	// Infrastructure
	logs := logger.GetLoggerFromCtx(ctx)

	db := postgres.New(ctx, &cfg.Config)

	hasher := hasher.New(cfg.LocalParameter)

	sender, err := smtp.NewSMTPSender(cfg.SMTPMail, cfg.SMTPPassword, cfg.SMTPHost, cfg.SMTPPort)
	if err != nil {
		logs.Fatal(ctx, "failed to create smtp sender", zap.Error(err))
	}

	// Repos, serivices and API handlers
	userRepository := repository.NewUserRepository(db)

	emailService := service.NewEmailService(sender, cfg.VerificationSubject, cfg.VerificationTemplate, cfg.Domain)
	userService := service.NewUserService(userRepository, hasher, emailService)

	handler := transport.NewHandler(userService, logs)

	// HTTP server
	srv := server.NewServer(cfg, handler.Init(cfg))

	go func() {
		if err := srv.Run(ctx); err != nil {
			logs.Error(ctx, "failed running http server", zap.Error(err))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	if err := srv.Stop(); err != nil {
		logs.Error(ctx, "failed shutting down the server", zap.Error(err))
	}

	db.Close()
}
