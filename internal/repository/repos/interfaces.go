package repos

import "app/pkg/domain/entity"

type Repositories struct {
	Auth IAuthRepo
	User IUserRepo
}

type IAuthRepo interface {
	Create(entity.UserAuth) error
}

type IUserRepo interface {
	FindByEmail(email string) (entity.User, error)
	CreatePersonal(newUser *entity.User) error
}
