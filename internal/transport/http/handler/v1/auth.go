package v1

import (
	"app/internal/service/dto"
	"app/pkg/api/response"
	"app/pkg/auth"
	"app/pkg/domain/entity"
	"app/pkg/infra/logger/sl"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"io"
	"log/slog"
	"net/http"
)

func (h *Handler) UserAuthTokenRefresh(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserSignIn"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	var body dto.UserRefreshTokenRequestDTO
	reqBodyError := render.DecodeJSON(r.Body, &body)
	if reqBodyError != nil && !errors.Is(reqBodyError, io.EOF) {
		log.Error("Failed to parse request body", sl.Err(reqBodyError))
		render.JSON(w, r, response.Error("Invalid params!"))
		return
	}

	isDataFromBody := !errors.Is(reqBodyError, io.EOF)

	if !isDataFromBody {
		preparedAuthId, err := r.Cookie(entity.AuthIdFieldName)
		if err != nil {
			log.Error("Request failed:", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("Bad request..."))
			return
		}

		authId, authIdErr := uuid.Parse(preparedAuthId.Value)
		if authIdErr != nil {
			log.Error("Request failed:", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, response.Error("Bad request..."))
			return
		}

		body.AuthId = authId
		refreshToken, err := r.Cookie(entity.RefreshTokenFieldName)
		if err == nil {
			body.RefreshToken = refreshToken.Value
		}

		accessToken, err := r.Cookie(entity.AccessTokenFieldName)
		if err == nil && accessToken.Value != "" {
			return
		}
	}

	if body.AuthId == uuid.Nil || body.RefreshToken == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error("Bad request! Fields authId, refreshToken is required!"))
		return
	}

	uAuth, err := h.services.Auth.RefreshTokens(body.AuthId, body.RefreshToken)
	if err != nil {
		log.Error("Failed to parse request body", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error("Refresh token is expired..."))
		return
	}
	auth.AuthorizeByCookieLevel(&uAuth.Token, uAuth.Id, w)

	render.JSON(w, r, dto.UserRefreshTokensResponseDTO{
		Status: response.StatusOK,
		Token:  &uAuth.Token,
		AuthId: uAuth.Id.String(),
	})
}
