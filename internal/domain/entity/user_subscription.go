package entity

import "github.com/google/uuid"

const (
	SubsResTypeUser     = "subs_res_type_user"
	SubsResTypeCategory = "subs_res_type_category"
)

type UserSubscription struct {
	SubscriberId uuid.UUID `json:"subscriber_id"`
	ResourceId   uuid.UUID `json:"resource_id"`
	ResourceType uuid.UUID `json:"resource_type"`
}
