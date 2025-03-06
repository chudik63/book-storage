package service

import (
	"book-storage/internal/models"
	"book-storage/pkg/email/smtp"
	"book-storage/pkg/logger"
	"context"
	"fmt"

	"go.uber.org/zap"
)

const (
	verificationLinkTmpl = "https://%s/verification?code=%s"
)

type verificationEmailInput struct {
	UserName         string
	VerificationLink string
}

type SMTPSender interface {
	Send(input smtp.SendEmailInput) error
}

type emailService struct {
	sender               SMTPSender
	verificationSubject  string
	verificationTemplate string

	domain string
}

func NewEmailService(s SMTPSender, vsub, vtempl, dom string) *emailService {
	return &emailService{
		sender:               s,
		verificationSubject:  vsub,
		verificationTemplate: vtempl,
		domain:               dom,
	}
}

func (s *emailService) SendVerificationEmail(ctx context.Context, inp *models.SendEmailInput) {
	logs := logger.GetLoggerFromCtx(ctx)

	templateInput := verificationEmailInput{
		UserName:         inp.UserName,
		VerificationLink: s.createVerificationLink(s.domain, inp.Code),
	}

	sendInput := smtp.SendEmailInput{Subject: s.verificationSubject, To: inp.Email}

	if err := sendInput.GenerateBodyFromHTML(s.verificationTemplate, templateInput); err != nil {
		logs.Error(ctx, "failed generate body from html template", zap.Error(err))

		return
	}

	err := s.sender.Send(sendInput)
	if err != nil {
		logs.Error(ctx, "failed send verification email", zap.Error(err))

		return
	}
}

func (s *emailService) createVerificationLink(domain, code string) string {
	return fmt.Sprintf(verificationLinkTmpl, domain, code)
}
