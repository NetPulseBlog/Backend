package entity

import "github.com/google/uuid"

// ResourceType
const (
	RTUser = iota + 1
	RTCategory
)

type UserSubscription struct {
	SubscriberId uuid.UUID `json:"subscriberId"`
	ResourceId   uuid.UUID `json:"subscribeId"`
	ResourceType uuid.UUID `json:"resourceType"`
}
