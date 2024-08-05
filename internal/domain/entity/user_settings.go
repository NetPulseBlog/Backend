package entity

import "github.com/google/uuid"

type NewsLineDefaultType string

const (
	NWDPopular NewsLineDefaultType = "nld_popular"
	NWDFresh   NewsLineDefaultType = "nld_fresh"
)

type NewsLineSortType string

const (
	NWSByPopular NewsLineSortType = "nls_by_popular"
	NWSByDate    NewsLineSortType = "nls_by_date"
)

type UserSettings struct {
	UserId          uuid.UUID           `json:"user_id"`
	NewsLineDefault NewsLineDefaultType `json:"news_line_default"`
	NewsLineSort    NewsLineSortType    `json:"news_line_sort"`
}
