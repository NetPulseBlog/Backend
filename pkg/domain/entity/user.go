package entity

import (
	"app/pkg/lib/ers"
	"app/pkg/lib/password"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserAccountType string
type UserRole string

const (
	UserAccountTypePersonal      UserAccountType = "personal"
	UserAccountTypeSystemSubSite UserAccountType = "system_sub_site"
)

const (
	UserRoleCustomer  UserRole = "customer"
	UserRoleAdmin     UserRole = "administrator"
	UserRoleModerator UserRole = "moderator"
)

type User struct {
	Id                uuid.UUID       `json:"id"`
	Email             string          `json:"email"`
	EncryptedPassword string          `json:"encrypted_password"`
	Salt              string          `json:"salt"`
	AccountType       UserAccountType `json:"account_type"`

	Role UserRole `json:"role"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	AvatarUrl string `json:"avatar_url"`
	CoverUrl  string `json:"cover_url"`

	Name        string `json:"name"`
	Description string `json:"description"`

	SubscribersCount   int `json:"subscribers_count"`
	SubscriptionsCount int `json:"subscriptions_count"`

	Settings UserSettings `json:"user_settings"`
}

func (u *User) CreatePassword(rawPassword string) error {
	const op = "entity.User.CreatePassword"

	salt, err := password.GenerateRandomSalt(password.DefaultSaltSize)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}
	u.Salt = salt

	passwordWithSalt := rawPassword + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	u.EncryptedPassword = string(hashedPassword)
	return nil
}
