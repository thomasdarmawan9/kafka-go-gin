package user_service

import (
	"users/domain/user_domain"
	"users/utils/error_utils"
)

var UserService userServiceInterface = &userService{}

type userServiceInterface interface {
	CreateUser(*user_domain.User) (*user_domain.User, error_utils.MessageErr)
	UserLogin(*user_domain.User) (string, error_utils.MessageErr)
}

type userService struct{}

func (u *userService) CreateUser(userReq *user_domain.User) (*user_domain.User, error_utils.MessageErr) {
	err := userReq.Validate()

	if err != nil {
		return nil, err
	}
	userReq.HashPass()

	res, err := user_domain.UserRepo.CreateUser(userReq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userService) UserLogin(userReq *user_domain.User) (string, error_utils.MessageErr) {
	userData, err := user_domain.UserRepo.GetUserByEmail(userReq.Email)
	if err != nil {
		return "", error_utils.NewNotAuthenticated("invalid email/password")
	}

	if ok := userReq.ComparePassword(userData.Password); !ok {
		return "", error_utils.NewNotAuthenticated("invalid email/password")
	}
	return userData.GenerateToken(), nil
}
