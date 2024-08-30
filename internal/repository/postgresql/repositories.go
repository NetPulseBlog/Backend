package postgresql

import (
	"app/internal/repository/repos"
	"database/sql"
)

func NewPostgresqlRepositories(db *sql.DB) *repos.Repositories {
	return &repos.Repositories{
		Auth:     NewAuthRepo(db),
		User:     NewUserRepo(db),
		Bookmark: NewBookmarkRepo(db),
		Article:  NewArticleRepo(db),
	}
}
