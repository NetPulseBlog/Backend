package service

import (
	"app/internal/repository/repos"
	"app/internal/service/dto"
	"app/pkg/domain/entity"
	"app/pkg/lib/ers"
	"github.com/google/uuid"
	"strings"
	"time"
)

type User struct {
	userRepo repos.IUserRepo

	authService *Auth
}

func NewUserService(userRepo repos.IUserRepo, authService *Auth) *User {
	return &User{userRepo: userRepo, authService: authService}
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
		return nil, ers.ThrowMessage(op, err)
	}

	if err := s.userRepo.CreatePersonal(&newUser); err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	return &newUser, nil
}

func (s *User) SignIn(reqData dto.UserSignInRequestDTO, deviceName string) (*entity.UserAuth, *entity.User, error) {
	const op = "service.User.SignIn"

	foundUser, err := s.userRepo.FindByEmail(reqData.Email)
	if err != nil {
		return nil, nil, ers.ThrowMessage(op, err)
	}

	if isValidPassword, err := foundUser.ComparePassword(reqData.Password); err != nil && !isValidPassword {
		return nil, nil, ers.ThrowMessage(op, entity.ErrUserInvalidPassword)
	}

	uAuth, err := s.authService.Authorize(foundUser, deviceName)
	if err != nil {
		return nil, nil, ers.ThrowMessage(op, err)
	}

	return uAuth, foundUser, nil
}
