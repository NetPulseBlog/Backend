package service

import (
	"app/internal/repository/repos"
	"app/pkg/domain/entity"
)

type Auth struct {
	authRepo repos.IAuthRepo
}

func NewAuthService(authRepo repos.IAuthRepo) *Auth {
	return &Auth{authRepo: authRepo}
}

func (s *Auth) Authorize(u *entity.User) (*entity.AuthToken, error) {
	return &entity.AuthToken{Access: "asd", Refresh: "ads"}, nil
}
