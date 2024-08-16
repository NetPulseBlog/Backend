package service

import (
	"app/internal/config"
	"app/internal/repository/repos"
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Auth struct {
	authRepo repos.IAuthRepo
	userRepo repos.IUserRepo
	cfg      *config.Config
}

func NewAuthService(authRepo repos.IAuthRepo, userRepo repos.IUserRepo, cfg *config.Config) *Auth {
	return &Auth{authRepo: authRepo, userRepo: userRepo, cfg: cfg}
}

func (s *Auth) VerifyToken(rawAuthId, accessToken string) error {
	const op = "service.Auth.VerifyToken"

	parsedToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.SecretKey), nil
	})
	if err != nil {
		return ers.ThrowMessage(op, fmt.Errorf("access token is invalid"))
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return ers.ThrowMessage(op, fmt.Errorf("access token is invalid"))
	}

	authId, err := uuid.Parse(rawAuthId)
	if err != nil {
		return err
	}
	uAuth, err := s.authRepo.GetById(authId)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	if uAuth.Token.RefreshExpiresAt.Unix() < time.Now().Unix() {
		return ers.ThrowMessage(op, fmt.Errorf("access token is invalid"))
	}

	return nil
}

func (s *Auth) RefreshTokens(authId uuid.UUID, refreshToken string) (*entity.UserAuth, error) {
	const op = "service.Auth.RefreshTokens"

	uAuth, err := s.authRepo.GetById(authId)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	if uAuth.Token.Refresh != refreshToken {
		return nil, ers.ThrowMessage(op, fmt.Errorf("invalid refresh token"))
	}

	parsedRefreshToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.SecretKey), nil
	})
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	_, ok := parsedRefreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ers.ThrowMessage(op, fmt.Errorf("invalid refresh token"))
	}

	if uAuth.Token.RefreshExpiresAt.Unix() < time.Now().Unix() {
		return nil, ers.ThrowMessage(op, fmt.Errorf("refresh token has expired"))
	}

	u, err := s.userRepo.FindById(uAuth.UserId)
	if err != nil {
		return nil, err
	}
	err = uAuth.GenerateTokens(u, s.cfg.JWT.SecretKey)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	err = s.authRepo.Update(uAuth)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	return uAuth, nil
}

func (s *Auth) Logout(authId uuid.UUID) error {
	const op = "service.Auth.Logout"
	if err := s.authRepo.DeleteItem(authId); err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}

func (s *Auth) Authorize(u *entity.User, deviceName string) (*entity.UserAuth, error) {
	const op = "service.Auth.Authorize"

	// find user auth by user_id and device_name
	// remove exist auth by device

	uAuth := entity.UserAuth{
		Id:     uuid.New(),
		UserId: u.Id,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		DeviceName: deviceName,
	}

	err := uAuth.GenerateTokens(u, s.cfg.SecretKey)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	if err := s.authRepo.Create(uAuth); err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	return &uAuth, nil
}
