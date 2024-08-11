package service

import (
	"app/internal/repository/repos"
)

type Deps struct {
	Repos *repos.Repositories
}

type Services struct {
	User User
	Auth Auth
}

func NewServices(deps Deps) *Services {
	userService := NewUserService(deps.Repos.User)
	authService := NewAuthService(deps.Repos.Auth)

	return &Services{
		User: *userService,
		Auth: *authService,
	}
}
