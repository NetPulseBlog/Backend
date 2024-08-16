package repos

import (
	"app/pkg/domain/entity"
	"github.com/google/uuid"
)

type Repositories struct {
	Auth     IAuthRepo
	User     IUserRepo
	Bookmark IBookmarkRepo
}

type IAuthRepo interface {
	GetById(uuid.UUID) (*entity.UserAuth, error)
	Update(*entity.UserAuth) error
	Create(entity.UserAuth) error
	DeleteItem(uuid.UUID) error
	RemoveExistsForDevice(uuid.UUID, string) error
}

type IUserRepo interface {
	FindById(uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	CreatePersonal(newUser *entity.User) error
}

type IBookmarkRepo interface {
	GetListByResourceType(resourceType entity.BookmarkResourceType) (*[]interface{}, error)
	Delete(id uuid.UUID) error
	Create(bookmark entity.UserBookmark) error
}
