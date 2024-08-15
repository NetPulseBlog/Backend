package v1

import (
	"app/internal/service/dto"
	"app/pkg/api/response"
	"app/pkg/auth"
	"app/pkg/domain/entity"
	"app/pkg/infra/logger/sl"
	vrules "app/pkg/lib/v-rules"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/mileusna/useragent"
	"log/slog"
	"net/http"
)

func (h *Handler) UserSignUp(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserSignUp"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var reqBody dto.UserSignUpRequestDTO
	err := render.DecodeJSON(r.Body, &reqBody)
	if err != nil {
		log.Error("Failed to parse request body", sl.Err(err))
		render.JSON(w, r, response.Error("Bad Request!"))
		return
	}

	validate := validator.New()
	validate.RegisterValidation("password", vrules.CustomPasswordValidation)

	if err := validate.Struct(reqBody); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		log.Error("Invalid request", sl.Err(err))
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	u, err := h.services.User.SignUp(reqBody)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.JSON(w, r, response.Error("Bad request..."))
		return
	}

	token, err := h.services.Auth.Authorize(u, auth.CreateDeviceNameFromUserAgent(useragent.Parse(r.UserAgent())))
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.JSON(w, r, response.Error("Bad request..."))
		return
	}

	auth.AuthorizeByCookieLevel(token, u.Id, w)

	render.JSON(w, r, dto.UserSignUpResponseDTO{
		Status: response.StatusOK,
		User:   dto.NewPublicUserResponseType(u),
		Token:  token,
	})
}

func (h *Handler) UserSignIn(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserSignIn"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var reqBody dto.UserSignInRequestDTO
	err := render.DecodeJSON(r.Body, &reqBody)
	if err != nil {
		log.Error("Failed to parse request body", sl.Err(err))
		render.JSON(w, r, response.Error("Bad Request!"))
		return
	}

	validate := validator.New()
	validate.RegisterValidation("password", vrules.CustomPasswordValidation)

	if err := validate.Struct(reqBody); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		log.Error("Invalid request", sl.Err(err))
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	t, u, err := h.services.User.SignIn(reqBody, auth.CreateDeviceNameFromUserAgent(useragent.Parse(r.UserAgent())))
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			log.Error("User not found:", sl.Err(err))
			render.JSON(w, r, response.Error("User not found..."))
			return
		}

		if errors.Is(err, entity.ErrUserInvalidPassword) {
			log.Error("User password is invalid:", sl.Err(err))
			render.JSON(w, r, response.Error("Password is invalid..."))
			return
		}

		log.Error("Request failed:", sl.Err(err))
		render.JSON(w, r, response.Error("Bad request..."))
		return
	}

	auth.AuthorizeByCookieLevel(t, u.Id, w)

	render.JSON(w, r, dto.UserSignUpResponseDTO{
		Status: response.StatusOK,
		User:   dto.NewPublicUserResponseType(u),
		Token:  t,
	})
}

func (h *Handler) UserSettingsUpdate(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserEdit(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserProfileByID(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserSubscribe(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserUnsubscribe(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserSubSites(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserPasswordRequestChange(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserPasswordConfirmChange(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
