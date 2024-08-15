package service

import (
	"app/internal/repository/repos"
	"app/internal/service/dto"
	"app/pkg/domain/entity"
	"github.com/google/uuid"
	"strings"
	"time"
)

type User struct {
	userRepo repos.IUserRepo
}

func NewUserService(userRepo repos.IUserRepo) *User {
	return &User{userRepo: userRepo}
}

func (s *User) SignUp(initialUserData dto.UserSignUpRequestDTO) (*entity.User, error) {
	const op = "service.User.SignUp"

	newUserId := uuid.New()
	newUserSettings := entity.UserSettings{
		UserId:          newUserId,
		NewsLineDefault: entity.NLDFresh,
		NewsLineSort:    entity.NLSByPopular,
	}
	newUser := entity.User{
		Id:          newUserId,
		Name:        strings.TrimSpace(initialUserData.Name),
		Description: "",
		Email:       strings.TrimSpace(initialUserData.Email),

		Role:        entity.UserRoleCustomer,
		AccountType: entity.UserAccountTypePersonal,

		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),

		AvatarUrl: "",
		CoverUrl:  "",

		Settings: newUserSettings,
	}

	err := newUser.CreatePassword(initialUserData.Password)
	if err != nil {
		return nil, err
	}

	if err := s.userRepo.CreatePersonal(&newUser); err != nil {
		return nil, err
	}

	return &newUser, nil
}
