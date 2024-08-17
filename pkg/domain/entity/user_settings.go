package entity

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

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

const NLDefaultValidationField = "nldefault"
const NLSortValidationField = "nlsort"

func NLDefaultValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	switch value {
	case
		string(NLDPopular),
		string(NLDFresh):
		return true
	}

	return true
}

func NLSortValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	switch value {
	case
		string(NLSByPopular),
		string(NLSByDate):
		return true
	}

	return false
}
