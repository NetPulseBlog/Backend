package entity

import "github.com/google/uuid"

type NewsLineDefaultType string

const (
	NLDPopular NewsLineDefaultType = "nld_popular"
	NLDFresh   NewsLineDefaultType = "nld_fresh"
)

type NewsLineSortType string

const (
	NLSByPopular NewsLineSortType = "nls_by_popular"
	NLSByDate    NewsLineSortType = "nls_by_date"
)

type UserSettings struct {
	UserId          uuid.UUID           `json:"user_id"`
	NewsLineDefault NewsLineDefaultType `json:"news_line_default"`
	NewsLineSort    NewsLineSortType    `json:"news_line_sort"`
}
