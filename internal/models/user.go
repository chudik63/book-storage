package models

import "time"

type User struct {
	ID               int64
	Login            string
	Name             string
	Email            string
	Password         string
	Salt             string
	VerificationCode string
	CreatedAt        time.Time
}
