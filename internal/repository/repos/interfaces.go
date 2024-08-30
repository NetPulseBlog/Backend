package repos

import (
	"app/pkg/domain/entity"
	"github.com/google/uuid"
)

type Repositories struct {
	Auth     IAuthRepo
	User     IUserRepo
	Bookmark IBookmarkRepo
	Article  IArticleRepo
}

type IArticleRepo interface {
	Create(article *entity.Article) error
	Update(article *entity.Article) error
	Delete(articleId string) error

	GetById(articleId string) (*entity.Article, error)
	GetList(listType string) (*[]entity.Article, error)
	ChangeStatus(articleId string) error

	CreateComment(comment *entity.ArticleComment) error
	GetCommentList(articleId string) (*[]entity.ArticleComment, error)
	UpdateComment(comment *entity.ArticleComment) error
	DeleteComment(commentId string) error
}

type IAuthRepo interface {
	GetById(uuid.UUID) (*entity.UserAuth, error)
	Update(*entity.UserAuth) error
	Create(entity.UserAuth) error
	DeleteItem(uuid.UUID) error
	RemoveExistsForDevice(uuid.UUID, string) error
}

type IUserRepo interface {
	GetSubSiteBarItems() (*[]entity.UserSubSiteBarItem, error)

	GetByAuthId(uuid.UUID) (*entity.User, error)

	FindById(uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	CreatePersonal(newUser *entity.User) error

	UpdateSettings(settings *entity.UserSettings) error

	Subscribe(subscription entity.UserSubscription) error
	Unsubscribe(ownerId, unsubscribedId uuid.UUID) error
}

type IBookmarkRepo interface {
	GetListByResourceType(userId uuid.UUID, resourceType entity.BookmarkResourceType) (*[]interface{}, error)
	Delete(id uuid.UUID) error
	Create(bookmark entity.UserBookmark) error
}
