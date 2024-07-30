package entity

import (
	"time"
)

type User struct {
	Id                int
	Email             string
	EncryptedPassword string

	CreatedAt time.Time
	UpdatedAt time.Time

	AvatarUrl string
	Name      string
	About     string

	Subscribers   int // get from virtual table based on user_subscription
	Subscriptions int // get from virtual table based on user_subscription

	Settings UserSettings
}
