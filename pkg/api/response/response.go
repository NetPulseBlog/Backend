package response

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string `json:"status"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	StatusOK    = "Oк"
	StatusError = "Error"
)

var (
	ErrValidationFailed    = errors.New("validation failed")
	ErrInternalServerError = errors.New("internal error")
	ErrUnknownId           = errors.New("unknown identifier")
	ErrBadRequest          = errors.New("bad request")
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	// todo: add json structure instead of string
	// todo: add password field
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid URL", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
