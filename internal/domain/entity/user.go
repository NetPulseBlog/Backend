package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id                uuid.UUID `json:"id"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"encryptedPassword"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	AvatarUrl string `json:"avatarUrl"`
	CoverUrl  string `json:"coverUrl"`

	Name        string `json:"name"`
	Description string `json:"description"`

	SubscribersCount   int `json:"subscribersCount"`
	SubscriptionsCount int `json:"subscriptionsCount"`

	Settings UserSettings `json:"userSettings"`
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
