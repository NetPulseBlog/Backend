package v1

import (
	"app/internal/service/dto"
	"app/pkg/api/request"
	"app/pkg/api/response"
	"app/pkg/auth"
	"app/pkg/domain/entity"
	"app/pkg/infra/logger/sl"
	vrules "app/pkg/lib/v-rules"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error(response.ErrBadRequest))
		return
	}

	validate := validator.New()
	validate.RegisterValidation("password", vrules.CustomPasswordValidation)

	if err := validate.Struct(reqBody); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		log.Error("Invalid request", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	u, err := h.services.User.SignUp(reqBody)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.JSON(w, r, response.Error(response.ErrBadRequest))
		return
	}

	uAuth, err := h.services.Auth.Authorize(u, auth.CreateDeviceNameFromUserAgent(useragent.Parse(r.UserAgent())))
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error(response.ErrBadRequest))
		return
	}

	auth.AuthorizeByCookieLevel(&uAuth.Token, uAuth.Id, w)

	render.JSON(w, r, dto.UserSignResponseDTO{
		Status: response.StatusOK,
		User:   dto.NewPublicUserResponseType(u),
		Token:  &uAuth.Token,
		AuthId: uAuth.Id.String(),
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
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error(response.ErrBadRequest))
		return
	}

	validate := validator.New()
	validate.RegisterValidation("password", vrules.CustomPasswordValidation)

	if err := validate.Struct(reqBody); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		log.Error("Invalid request", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.ValidationError(validateErr))
		return
	}

	uAuth, u, err := h.services.User.SignIn(reqBody, auth.CreateDeviceNameFromUserAgent(useragent.Parse(r.UserAgent())))
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			log.Error("User not found:", sl.Err(err))
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, response.Error(errors.New("User not found...")))
			return
		}

		if errors.Is(err, entity.ErrUserInvalidPassword) {
			log.Error("User password is invalid:", sl.Err(err))
			render.Status(r, http.StatusMethodNotAllowed)
			render.JSON(w, r, response.Error(errors.New("Password is invalid...")))
			return
		}

		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error(response.ErrBadRequest))
		return
	}

	auth.AuthorizeByCookieLevel(&uAuth.Token, uAuth.Id, w)

	render.JSON(w, r, dto.UserSignResponseDTO{
		Status: response.StatusOK,
		User:   dto.NewPublicUserResponseType(u),
		Token:  &uAuth.Token,
		AuthId: uAuth.Id.String(),
	})
}

func (h *Handler) UserSettingsUpdate(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UserEdit(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserProfileByID(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserProfileByID"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	u, err := h.services.User.GetUser(userId)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.JSON(w, r, dto.NewPublicUserResponseType(u))
}

func (h *Handler) UserProfile(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserProfile"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	authId, err := request.GetAuthId(r)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	user, err := h.services.User.GetUserByAuthId(authId)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.JSON(w, r, dto.NewPublicUserResponseType(user))
}

func (h *Handler) UserSubscribe(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserSubscribe"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	subscribedId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	authId, err := request.GetAuthId(r)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	user, err := h.services.User.GetUserByAuthId(authId)

	err = h.services.User.Subscribe(user.Id, subscribedId)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.JSON(w, r, response.OK())
}

func (h *Handler) UserUnsubscribe(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserUnsubscribe"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	unsubscribedId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error(response.ErrBadRequest))
		return
	}

	authId, err := request.GetAuthId(r)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, response.Error(response.ErrBadRequest))
		return
	}

	user, err := h.services.User.GetUserByAuthId(authId)

	err = h.services.User.Unsubscribe(user.Id, unsubscribedId)
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.JSON(w, r, response.OK())
}

func (h *Handler) UserSubSites(w http.ResponseWriter, r *http.Request) {
	const op = "http.v1.User.UserSubSites"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	subSites, err := h.services.User.GetSubSiteBarItems()
	if err != nil {
		log.Error("Request failed:", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.Error(response.ErrInternalServerError))
		return
	}

	render.JSON(w, r, subSites)
}

func (h *Handler) UserPasswordRequestChange(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func (h *Handler) UserPasswordConfirmChange(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
