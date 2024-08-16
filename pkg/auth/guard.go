package auth

import (
	"app/internal/service"
	"app/pkg/domain/entity"
	"net/http"
)

func CreateGuardMiddleware(authService service.Auth) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken, _ := r.Cookie(entity.AccessTokenFieldName)
			if err := accessToken.Valid(); err != nil {
				http.Error(w, "Access token is missing", http.StatusUnauthorized)
				return
			}

			authId, _ := r.Cookie(entity.AuthIdFieldName)
			if err := authId.Valid(); err != nil {
				http.Error(w, "AuthId is missing", http.StatusUnauthorized)
				return
			}

			err := authService.VerifyToken(authId.Value, accessToken.Value)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
