package entity

import (
	"github.com/google/uuid"
	"time"
)

type UserSubscription struct {
	OwnerId          uuid.UUID `json:"owner_id"`
	SubscribedUserId uuid.UUID `json:"subscribed_user_id"`

	CreatedAt time.Time `json:"created_at"`
}
