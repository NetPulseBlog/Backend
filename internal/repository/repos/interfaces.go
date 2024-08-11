package repos

import "app/pkg/domain/entity"

type Repositories struct {
	Auth IAuthRepo
	User IUserRepo
}

type IAuthRepo interface {
}

type IUserRepo interface {
	CreatePersonal(newUser *entity.User) error
}
