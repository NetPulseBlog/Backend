package entity

import (
	"app/pkg/lib/ers"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const (
	AccessTokenFieldName  = "access_token"
	RefreshTokenFieldName = "refresh_token"
	AuthIdFieldName       = "auth_id"
)

const (
	JWTAccessExpiresDaysInHours  = 24 * 7  // 7 days
	JWTRefreshExpiresDaysInHours = 24 * 30 // 1 month
)

var (
	ErrAccessTokenIsInvalid = errors.New("access token is invalid")
	ErrAccessTokenIsExpired = errors.New("access token is expired")
)

type JwtClaims struct {
	UserId uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}

type UserAuth struct {
	Id     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeviceName string `json:"device_name"`
	Token      AuthToken

	// todo Last Activity by current auth
}

func (a *UserAuth) GenerateTokens(u *User, SecretKey string) error {
	const op = "entity.Auth.GenerateTokens"

	uaToken := AuthToken{
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
	accessTokenString, err := accessToken.SignedString([]byte(SecretKey))
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(SecretKey))
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	uaToken.Access = accessTokenString
	uaToken.Refresh = refreshTokenString

	a.Token = uaToken

	return nil
}

type AuthToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`

	AccessExpiresAt  time.Time `json:"access_expires_at"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
}
