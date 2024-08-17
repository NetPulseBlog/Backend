package entity

import (
	"app/pkg/lib/ers"
	"app/pkg/lib/password"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserRole string

const (
	UserRoleCustomer  UserRole = "customer"
	UserRoleAdmin     UserRole = "administrator"
	UserRoleModerator UserRole = "moderator"
)

var (
	ErrUserNotFound        = errors.New("User not found")
	ErrUserInvalidPassword = errors.New("User password is invalid")
)

type User struct {
	Id                uuid.UUID `json:"id"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"encrypted_password"`
	Salt              string    `json:"salt"`

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

type UserSubSiteBarItem struct {
	Id        uuid.UUID `json:"id"`
	AvatarUrl string    `json:"avatar_url"`
	Name      string    `json:"name"`
}

func (u *User) CreatePassword(rawPassword string) error {
	const op = "entity.User.CreatePassword"

	salt, err := password.GenerateRandomSalt(password.DefaultSaltSize)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}
	u.Salt = salt

	passwordWithSalt := strings.TrimSpace(rawPassword) + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	u.EncryptedPassword = string(hashedPassword)
	fmt.Println(u.EncryptedPassword)

	return nil
}

func (u *User) ComparePassword(rawPassword string) (bool, error) {
	const op = "entity.User.ComparePassword"

	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(rawPassword+u.Salt))
	return err == nil, ers.ThrowMessage(op, err)
}
