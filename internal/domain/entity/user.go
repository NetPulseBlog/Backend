package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id                uuid.UUID `json:"id"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"encrypted_password"`

	// TODO: Hash, ConfirmCode, Role, Phone

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

//// Validate ...
//func (u *User) Validate() error {
//	return validation.ValidateStruct(
//		u,
//		validation.Field(&u.Email, validation.Required, is.Email),
//		validation.Field(&u.Password,
//			validation.By(requiredIf(u.EncryptedPassword == "")),
//			validation.Length(6, 100),
//		),
//	)
//}
//
//// BeforeCreate ...
//func (u *User) BeforeCreate() error {
//	if len(u.Password) > 0 {
//		enc, err := encryptString(u.Password)
//
//		if err != nil {
//			return err
//		}
//
//		u.EncryptedPassword = enc
//	}
//
//	return nil
//}
//
//func encryptString(s string) (string, error) {
//	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
//	if err != nil {
//		return "", err
//	}
//
//	return string(b), nil
//}
