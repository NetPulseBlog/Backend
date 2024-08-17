package dto

import (
	"app/pkg/domain/entity"
	"github.com/google/uuid"
	"time"
)

type UserRefreshTokenRequestDTO struct {
	AuthId       uuid.UUID `json:"email"`
	RefreshToken string    `json:"refreshToken"`
	AccessToken  string    `json:"accessToken"`
}

type UserSignInRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type UserSignUpRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type PublicUserResponseType struct {
	Id                 uuid.UUID           `json:"id"`
	Name               string              `json:"name"`
	Email              string              `json:"email"`
	Role               entity.UserRole     `json:"role"`
	CreatedAt          time.Time           `json:"createdAt"`
	UpdatedAt          time.Time           `json:"updatedAt"`
	AvatarUrl          string              `json:"avatarUrl"`
	CoverUrl           string              `json:"coverUrl"`
	Description        string              `json:"description"`
	SubscribersCount   int                 `json:"subscribersCount"`
	SubscriptionsCount int                 `json:"subscriptionsCount"`
	Settings           entity.UserSettings `json:"settings"`
}

func NewPublicUserResponseType(u *entity.User) *PublicUserResponseType {
	pu := PublicUserResponseType{}

	pu.Id = u.Id
	pu.Name = u.Name
	pu.Email = u.Email
	pu.Role = u.Role
	pu.CreatedAt = u.CreatedAt
	pu.UpdatedAt = u.UpdatedAt
	pu.AvatarUrl = u.AvatarUrl
	pu.CoverUrl = u.CoverUrl
	pu.Description = u.Description
	pu.SubscribersCount = u.SubscribersCount
	pu.SubscriptionsCount = u.SubscriptionsCount
	pu.Settings = u.Settings

	return &pu
}

type UserSignResponseDTO struct {
	Status string                  `json:"status"`
	User   *PublicUserResponseType `json:"user"`
	Token  *entity.AuthToken       `json:"token"`
	AuthId string                  `json:"authId"`
}

type UserRefreshTokensResponseDTO struct {
	Status string            `json:"status"`
	Token  *entity.AuthToken `json:"token"`
	AuthId string            `json:"authId"`
}
