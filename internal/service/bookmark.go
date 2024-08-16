package service

import (
	"app/internal/repository/repos"
	"app/pkg/domain/entity"
)

type Bookmark struct {
	bookmarkRepo repos.IBookmarkRepo
}

func NewBookmarkService(bookmarkRepo repos.IBookmarkRepo) *Bookmark {
	return &Bookmark{bookmarkRepo: bookmarkRepo}
}

func (s Bookmark) GetList(resourceType entity.BookmarkResourceType) (*[]interface{}, error) {
	return s.bookmarkRepo.GetListByResourceType(resourceType)
}

func (s Bookmark) Create() {
	// todo
}

func (s Bookmark) Delete() {
	// todo
}
