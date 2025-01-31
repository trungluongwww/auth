package usecase

import (
	"context"
	"errors"
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
	"github.com/trungluongwww/auth/pkg/repository"
	"github.com/trungluongwww/auth/pkg/service"
)

type User interface {
	Register(context context.Context, p request.RegisterPayload) error
	Login(context context.Context, p request.LoginPayload) (*response.LoginResponse, error)
	LoginWithFacebook(context context.Context) error
	LoginWithGoogle(context context.Context) error
	GetMe(context context.Context, id int) (*response.UserResponse, error)
}

type userImpl struct {
	Repository  repository.Repository
	AuthService service.AuthService
	UserService service.UserService
}

func NewUser(repository repository.Repository, authService service.AuthService, userService service.UserService) User {
	return &userImpl{
		Repository:  repository,
		AuthService: authService,
		UserService: userService,
	}
}

func (u *userImpl) Register(context context.Context, p request.RegisterPayload) error {
	exist, _ := u.Repository.NewUser().FirstRaw(&model.User{Email: p.Email})
	if exist != nil {
		return errors.New("user already exist")
	}

	err := u.Repository.NewTransaction(func(tx repository.Repository) error {
		var (
			userDao = tx.NewUser()
		)
		account, err := u.AuthService.ConvertToAccountModel(p.Password)
		if err != nil {
			return err
		}

		err = userDao.InsertAccount(account)
		if err != nil {
			return err
		}

		user := u.UserService.ConvertRegisterPayloadToModel(p, account)
		err = userDao.Insert(user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *userImpl) Login(context context.Context, p request.LoginPayload) (*response.LoginResponse, error) {
	exist, err := u.Repository.NewUser().FirstRaw(&model.User{Email: p.Email})
	if err != nil {
		return nil, err
	}

	if exist.IsBanned {
		return nil, errors.New("user is banned")
	}

	account, err := u.Repository.NewUser().FirstAccountRaw(&model.Account{ID: exist.AccountID})
	if err != nil {
		return nil, err
	}

	accessTokenRes, rf, err := u.AuthService.GenerateAccessToken(account, exist.Email)
	if err != nil {
		return nil, err
	}

	err = u.Repository.NewUser().InsertAccountRefreshToken(rf)
	if err != nil {
		return nil, err
	}

	userRes := u.UserService.ConvertToUserResponse(exist)

	return &response.LoginResponse{
		User:           *userRes,
		AccessResponse: *accessTokenRes,
	}, nil
}

func (u *userImpl) LoginWithFacebook(context context.Context) error {
	return nil
}

func (u *userImpl) LoginWithGoogle(context context.Context) error {
	return nil
}

func (u *userImpl) GetMe(context context.Context, id int) (*response.UserResponse, error) {
	exist, err := u.Repository.NewUser().FirstRaw(&model.User{ID: uint32(id)})
	if err != nil {
		return nil, err
	}

	return u.UserService.ConvertToUserResponse(exist), nil
}
