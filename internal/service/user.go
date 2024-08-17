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
	authRepo repos.IAuthRepo

	authService *Auth
}

func NewUserService(userRepo repos.IUserRepo, authRepo repos.IAuthRepo, authService *Auth) *User {
	return &User{userRepo: userRepo, authRepo: authRepo, authService: authService}
}

func (s *User) GetUser(userId uuid.UUID) (*entity.User, error) {
	const op = "service.User.GetUser"

	u, err := s.userRepo.FindById(userId)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	return u, nil
}

func (s *User) GetSubSiteBarItems() (*[]entity.UserSubSiteBarItem, error) {
	const op = "service.User.GetSubSiteBarItems"

	items, err := s.userRepo.GetSubSiteBarItems()
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	return items, nil
}

func (s *User) GetUserByAuthId(authId uuid.UUID) (*entity.User, error) {
	const op = "service.User.GetUserByAuthId"

	u, err := s.userRepo.GetByAuthId(authId)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	return u, nil
}

func (s *User) UpdateSettings(reqBody dto.UpdateUserSettingsRequestDTO, userId uuid.UUID) (*entity.UserSettings, error) {
	const op = "service.User.UpdateSettings"

	us := entity.UserSettings{
		UserId:          userId,
		NewsLineDefault: reqBody.NewsLineDefault,
		NewsLineSort:    reqBody.NewsLineSort,
	}

	err := s.userRepo.UpdateSettings(&us)
	if err != nil {
		return nil, ers.ThrowMessage(op, err)
	}

	return &us, nil
}

func (s *User) Subscribe(ownerId, subscribedId uuid.UUID) error {
	const op = "service.User.Subscribe"

	uSub := entity.UserSubscription{
		OwnerId:          ownerId,
		SubscribedUserId: subscribedId,
		CreatedAt:        time.Now().UTC(),
	}

	err := s.userRepo.Subscribe(uSub)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
}

func (s *User) Unsubscribe(ownerId, unsubscribedId uuid.UUID) error {
	const op = "service.User.Unsubscribe"

	err := s.userRepo.Unsubscribe(ownerId, unsubscribedId)
	if err != nil {
		return ers.ThrowMessage(op, err)
	}

	return nil
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

		Role: entity.UserRoleCustomer,

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
