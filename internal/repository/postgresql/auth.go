package postgresql

import "database/sql"

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

//func (receiver ExampleRepo) Create(name string) string {
//	return "Created"
//}
//
//func (receiver ExampleRepo) Update() {
//	panic("IMPLEMENT ME...")
//}
//
//func (receiver ExampleRepo) GetList() {
//	panic("IMPLEMENT ME...")
//}
//
//func (receiver ExampleRepo) Delete() {
//	panic("IMPLEMENT ME...")
//}
