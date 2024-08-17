package service

import (
	"app/internal/repository/repos"
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"github.com/google/uuid"
	"time"
)

type Bookmark struct {
	bookmarkRepo repos.IBookmarkRepo
	authRepo     repos.IAuthRepo
}

func NewBookmarkService(bookmarkRepo repos.IBookmarkRepo, authRepo repos.IAuthRepo) *Bookmark {
	return &Bookmark{bookmarkRepo: bookmarkRepo, authRepo: authRepo}
}

func (s Bookmark) GetList(userId uuid.UUID, resourceType entity.BookmarkResourceType) (*[]interface{}, error) {
	return s.bookmarkRepo.GetListByResourceType(userId, resourceType)
}

func (s Bookmark) Create(authId uuid.UUID, resourceId string, resourceType entity.BookmarkResourceType) error {
	const op = "service.Bookmark.Create"

	resourceUuid, err := uuid.Parse(resourceId)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	uAuth, err := s.authRepo.GetById(authId)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	bookmark := entity.UserBookmark{
		UserId:       uAuth.UserId,
		ResourceId:   resourceUuid,
		ResourceType: resourceType,
		CreatedAt:    time.Now().UTC(),
	}

	err = s.bookmarkRepo.Create(bookmark)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}

func (s Bookmark) Delete(resourceId string) error {
	const op = "service.Bookmark.Delete"

	resourceUuid, err := uuid.Parse(resourceId)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	err = s.bookmarkRepo.Delete(resourceUuid)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}
