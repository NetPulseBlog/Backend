package entity

import "github.com/google/uuid"

// NewsLineDefault
const (
	NWDPopular = "nld_popular"
	NWDFresh   = "nld_fresh"
)

// NewsLineSort
const (
	NWSByPopular = "nls_by_popular"
	NWSByDate    = "nls_by_date"
)

type UserSettings struct {
	UserId          uuid.UUID `json:"user_id"`
	NewsLineDefault string    `json:"news_line_default"`
	NewsLineSort    string    `json:"news_line_sort"`
}
