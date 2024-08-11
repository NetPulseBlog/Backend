package auth

import (
	"app/pkg/domain/entity"
	"net/http"
)

const (
	AccessTokenFieldName  = "access_token"
	RefreshTokenFieldName = "refresh_token"
)

func AuthorizeByCookieLevel(token *entity.AuthToken, w http.ResponseWriter) {
	// write tokens to
}
