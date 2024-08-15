package service

import (
	"app/internal/config"
	"app/internal/repository/repos"
	"app/pkg/domain/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Auth struct {
	authRepo repos.IAuthRepo
	cfg      *config.Config
}

const (
	JWTAccessExpiresDaysInHours  = 24 * 7  // 7 days
	JWTRefreshExpiresDaysInHours = 24 * 30 // 1 month
)

type JwtClaims struct {
	UserId uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}

func NewAuthService(authRepo repos.IAuthRepo, cfg *config.Config) *Auth {
	return &Auth{authRepo: authRepo, cfg: cfg}
}

func (s *Auth) Authorize(u *entity.User, deviceName string) (*entity.AuthToken, error) {
	uAuth := entity.UserAuth{
		Id:     uuid.New(),
		UserId: u.Id,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		DeviceName: deviceName,
	}
	uaToken := entity.AuthToken{
		RefreshExpiresAt: time.Now().Add(time.Hour * JWTRefreshExpiresDaysInHours),
		AccessExpiresAt:  time.Now().Add(time.Hour * JWTAccessExpiresDaysInHours),
	}

	accessTokenClaims := &JwtClaims{
		UserId: u.Id,
		Email:  u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: uaToken.AccessExpiresAt.Unix(),
		},
	}

	refreshTokenClaims := &JwtClaims{
		UserId: u.Id,
		Email:  u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: uaToken.RefreshExpiresAt.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.cfg.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.cfg.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	uaToken.Access = accessTokenString
	uaToken.Refresh = refreshTokenString

	uAuth.Token = uaToken

	if err := s.authRepo.Create(uAuth); err != nil {
		return nil, err
	}

	return &uAuth.Token, nil
}
