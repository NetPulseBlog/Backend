package postgresql

import (
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type BookmarkRepo struct {
	db *sql.DB
}

func NewBookmarkRepo(db *sql.DB) *BookmarkRepo {
	return &BookmarkRepo{
		db: db,
	}
}

func (repo BookmarkRepo) GetListByResourceType(resourceType entity.BookmarkResourceType) (*[]interface{}, error) {
	const op = "postgresql.BookmarkRepo.GetById"

	sqlQuery := ""

	if resourceType == entity.BTComment {
		sqlQuery = `
			SELECT
				ac.*
			FROM
				user_bookmark ub
			JOIN article_comment ac ON ac.id = ub.resource_id
			WHERE
				ub.resource_type = $1
		`
	} else if resourceType == entity.BTArticle {
		sqlQuery = `
			SELECT
				a.*
			FROM
				user_bookmark ub
			JOIN article a ON a.id = ub.resource_id
			WHERE
				ub.resource_type = $1
		`
	}

	if sqlQuery == "" {
		return nil, ers.ThrowMessage(op, fmt.Errorf("empty query"))
	}

	rows, err := repo.db.Query(sqlQuery, resourceType)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	defer rows.Close()

	var bList []interface{}

	for rows.Next() {
		var row interface{}
		err = rows.Scan(&row)

		if err != nil {
			log.Println(err)
			return nil, ers.ThrowMessage(op, err)
		}

		bList = append(bList, row)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	return &bList, nil
}

func (repo BookmarkRepo) Delete(id uuid.UUID) error {
	return nil
}

func (repo BookmarkRepo) Create(resourceId uuid.UUID, resourceType entity.BookmarkResourceType) error {
	return nil
}
