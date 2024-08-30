package service

import (
	"app/internal/config"
	"app/internal/repository/repos"
)

type Deps struct {
	Repos  *repos.Repositories
	Config *config.Config
}

type Services struct {
	User     User
	Auth     Auth
	Bookmark Bookmark
	Article  Article
}

func NewServices(deps Deps) *Services {
	authService := NewAuthService(deps.Repos.Auth, deps.Repos.User, deps.Config)
	userService := NewUserService(deps.Repos.User, deps.Repos.Auth, authService)
	articleService := NewArticleService(deps.Repos.Article, deps.Repos.User, deps.Config)
	bookmarkService := NewBookmarkService(deps.Repos.Bookmark, deps.Repos.Auth)

	return &Services{
		User:     *userService,
		Auth:     *authService,
		Bookmark: *bookmarkService,
		Article:  *articleService,
	}
}
