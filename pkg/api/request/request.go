package request

import (
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"errors"
	"github.com/google/uuid"
	"net/http"
)

func GetAuthId(r *http.Request) (uuid.UUID, error) {
	const op = "request.GetAuthId"

	var authId string

	authIdHeader := r.Header.Get(entity.AccessTokenHeaderFieldName)
	if authIdHeader == "" {
		authIdCookie, _ := r.Cookie(entity.AuthIdFieldName)
		if err := authIdCookie.Valid(); err != nil {
			return uuid.Nil, err
		}

		authId = authIdCookie.Value
	} else {
		authId = authIdHeader
	}

	if authId == "" {
		return uuid.Nil, ers.ThrowMessage(op, errors.New("auth id not found"))
	}

	authUuid, err := uuid.Parse(authId)
	if err != nil {
		return uuid.Nil, ers.ThrowMessage(op, err)
	}

	return authUuid, nil
}
