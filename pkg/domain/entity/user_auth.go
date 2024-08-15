package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	AccessTokenFieldName  = "access_token"
	RefreshTokenFieldName = "refresh_token"
)

type UserAuth struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeviceName string `json:"device_name"`
	Token      AuthToken

	// todo Last Activity by current auth
}

type AuthToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`

	AccessExpiresAt  time.Time `json:"access_expires_at"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
}
