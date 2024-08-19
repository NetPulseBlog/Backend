package service

import (
	"app/internal/config"
	"app/internal/repository/repos"
	"app/internal/service/dto"
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"github.com/google/uuid"
	"time"
)

type Article struct {
	articleRepo repos.IArticleRepo
	userRepo    repos.IUserRepo
	cfg         *config.Config
}

func NewArticleService(articleRepo repos.IArticleRepo, userRepo repos.IUserRepo, cfg *config.Config) *Article {
	return &Article{articleRepo: articleRepo, userRepo: userRepo, cfg: cfg}
}

func (s Article) Create(user *entity.User, rawArticle *dto.CreateArticleRequestDTO) error {
	const op = "service.Article.Create"

	newArticle := entity.Article{
		Id:       uuid.New(),
		AuthorId: user.Id,

		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),

		Status: entity.ArticleStatus(rawArticle.Status),

		Title:      rawArticle.Title,
		SubsSiteId: rawArticle.SubsSiteId,

		/*
			TODO:
				ContentBlocks
				CoverUrl
				SubTitle
		*/

		ViewsCount: 0,
	}

	// fill cover url and subtitle

	err := s.articleRepo.Create(&newArticle)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	// todo: notify subscribers if publish

	return nil
}

func (s Article) Update() {
	const op = "service.Article.Update"

}

func (s Article) Delete(id string) error {
	const op = "service.Article.Delete"

	err := s.articleRepo.Delete(id)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}

func (s Article) GetById() {
	const op = "service.Article.GetById"

}

func (s Article) GetList() {
	const op = "service.Article.GetList"

}

func (s Article) ChangeStatus() {
	const op = "service.Article.ChangeStatus"
	// todo: notify subscribers if publish
}

func (s Article) CreateComment() {
	const op = "service.Article.CreateComment"
	// todo: notify subscribers
}

func (s Article) GetCommentList() {
	const op = "service.Article.GetCommentList"

}

func (s Article) UpdateComment() {
	const op = "service.Article.UpdateComment"

}

func (s Article) DeleteComment() {
	const op = "service.Article.DeleteComment"

}
