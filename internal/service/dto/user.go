package dto

import (
	"app/pkg/domain/entity"
	"time"
)

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
	Name               string                 `json:"name"`
	Email              string                 `json:"email"`
	AccountType        entity.UserAccountType `json:"account_type"`
	Role               entity.UserRole        `json:"role"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
	AvatarUrl          string                 `json:"avatar_url"`
	CoverUrl           string                 `json:"cover_url"`
	Description        string                 `json:"description"`
	SubscribersCount   int                    `json:"subscribers_count"`
	SubscriptionsCount int                    `json:"subscriptions_count"`
	Settings           entity.UserSettings    `json:"user_settings"`
}

func NewPublicUserResponseType(u *entity.User) *PublicUserResponseType {
	pu := PublicUserResponseType{}

	pu.Name = u.Name
	pu.Email = u.Email
	pu.AccountType = u.AccountType
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

type UserSignUpResponseDTO struct {
	Status string                  `json:"status"`
	User   *PublicUserResponseType `json:"user"`
	Token  *entity.AuthToken       `json:"token"`
}
