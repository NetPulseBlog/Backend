package entity

import "time"

/*
TODO: add struct tags
*/

type ArticleComment struct {
	Id        int
	UserId    int
	ArticleId int
	ReplyId   int

	CreatedAt time.Time

	Content string
}
