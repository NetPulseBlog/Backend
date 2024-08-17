package service

import (
	"app/internal/config"
	"app/internal/repository/repos"
)

type Article struct {
	articleRepo repos.IArticleRepo
	userRepo    repos.IUserRepo
	cfg         *config.Config
}

func NewArticleService(articleRepo repos.IArticleRepo, userRepo repos.IUserRepo, cfg *config.Config) *Article {
	return &Article{articleRepo: articleRepo, userRepo: userRepo, cfg: cfg}
}
