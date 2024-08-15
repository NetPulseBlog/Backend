package auth

import (
	"app/pkg/domain/entity"
	"github.com/google/uuid"
	"github.com/mileusna/useragent"
	"net/http"
)

func CreateDeviceNameFromUserAgent(ua useragent.UserAgent) string {
	return ua.Name + " " + ua.OS + "(" + ua.String + ")"
}

func AuthorizeByCookieLevel(token *entity.AuthToken, authId uuid.UUID, w http.ResponseWriter) {
	accessCookie := &http.Cookie{
		Name:     entity.AccessTokenFieldName,
		Value:    token.Access,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  token.AccessExpiresAt,
	}

	refreshCookie := &http.Cookie{
		Name:     entity.RefreshTokenFieldName,
		Value:    token.Refresh,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  token.RefreshExpiresAt,
	}

	authIdCookie := &http.Cookie{
		Name:     entity.AuthIdFieldName,
		Value:    authId.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	// Записываем куки в ответ
	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, authIdCookie)
}
