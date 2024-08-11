package entity

import (
	"github.com/google/uuid"
	"time"
)

type UserAuth struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`

	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`

	DeviceName string `json:"device_name"`

	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
