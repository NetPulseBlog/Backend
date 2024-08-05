package entity

import (
	"github.com/google/uuid"
	"time"
)

type UserAuth struct {
	Id     uuid.UUID `json:"user_id"`
	UserId uuid.UUID `json:"user_id"`

	refreshToken string `json:"refresh_token"`
	accessToken  string `json:"access_token"`

	deviceName string `json:"device_name"`

	expiresAt time.Time `json:"expires_at"`
	createdAt time.Time `json:"created_at"`
}
