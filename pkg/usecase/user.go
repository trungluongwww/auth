package usecase

import (
	"context"
	"errors"
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
	"github.com/trungluongwww/auth/pkg/repository"
	"github.com/trungluongwww/auth/pkg/service"
	"github.com/trungluongwww/auth/third_party/social"
)

type User interface {
	Register(context context.Context, p request.RegisterPayload) error
	Login(context context.Context, p request.LoginPayload) (*response.LoginResponse, error)
	RefreshToken(context context.Context, p *request.RefreshTokenPayload) (*response.LoginResponse, error)
	LoginWithFacebook(context context.Context, p request.FacebookLoginPayload) (*response.LoginResponse, error)
	GetMe(context context.Context, id int) (*response.UserResponse, error)
}

type userImpl struct {
	Repository  repository.Repository
	AuthService service.AuthService
	UserService service.UserService
	Social      social.Social
}

func NewUser(repository repository.Repository, authService service.AuthService, userService service.UserService,
	social social.Social) User {
	return &userImpl{
		Repository:  repository,
		AuthService: authService,
		UserService: userService,
		Social:      social,
	}
}

func (u *userImpl) Register(context context.Context, p request.RegisterPayload) error {
	user, _ := u.Repository.NewUser().FirstRaw(&model.User{Email: p.Email})
	if user != nil {
		return errors.New("user already user")
	}

	err := u.Repository.NewTransaction(func(tx repository.Repository) error {
		var (
			userDao    = tx.NewUser()
			accountDao = tx.NewAccount()
		)
		account, err := u.AuthService.ConvertToAccountModel(p.Password)
		if err != nil {
			return err
		}

		err = accountDao.Insert(account)
		if err != nil {
			return err
		}

		user := u.UserService.ConvertRegisterPayloadToModel(p, account, false)
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
	var (
		accountDao = u.Repository.NewAccount()
	)
	user, err := u.Repository.NewUser().FirstRaw(&model.User{Email: p.Email})
	if err != nil {
		return nil, err
	}

	if user.IsBanned {
		return nil, errors.New("user is banned")
	}

	account, err := accountDao.FirstRaw(&model.Account{ID: user.AccountID})
	if err != nil {
		return nil, err
	}

	accessTokenRes, rf, err := u.AuthService.GenerateAccessToken(account.ID, user.Email)
	if err != nil {
		return nil, err
	}

	err = accountDao.InsertAccountRefreshToken(rf)
	if err != nil {
		return nil, err
	}

	userRes := u.UserService.ConvertToUserResponse(user)

	return &response.LoginResponse{
		User:           *userRes,
		AccessResponse: *accessTokenRes,
	}, nil
}

func (u *userImpl) LoginWithFacebook(context context.Context, p request.FacebookLoginPayload) (*response.LoginResponse, error) {
	fbData, err := u.Social.GetFacebookInfo(p.Token)
	if err != nil {
		return nil, err
	}

	user, err := u.Repository.NewUser().FirstRaw(&model.User{Email: fbData.Email})
	if err != nil {
		return nil, err
	}

	res := &response.LoginResponse{}

	if user != nil {
		accessTokenRes, rf, err := u.AuthService.GenerateAccessToken(user.AccountID, user.Email)
		if err != nil {
			return nil, err
		}

		err = u.Repository.NewAccount().InsertAccountRefreshToken(rf)
		if err != nil {
			return nil, err
		}

		res.User = *u.UserService.ConvertToUserResponse(user)
		res.AccessResponse = *accessTokenRes
	} else {
		err = u.Repository.NewTransaction(func(tx repository.Repository) error {
			var (
				userDao    = tx.NewUser()
				accountDao = tx.NewAccount()
			)

			account, err := u.AuthService.ConvertToAccountModel("")
			if err != nil {
				return err
			}

			err = accountDao.Insert(account)
			if err != nil {
				return err
			}

			payload := request.RegisterPayload{
				Email: fbData.Email,
				Name:  fbData.Name,
			}

			newUser := u.UserService.ConvertRegisterPayloadToModel(payload, account, true)
			err = userDao.Insert(newUser)
			if err != nil {
				return err
			}

			accessTokenRes, rf, err := u.AuthService.GenerateAccessToken(account.ID, newUser.Email)
			if err != nil {
				return err
			}

			err = accountDao.InsertAccountRefreshToken(rf)
			if err != nil {
				return err
			}

			res.User = *u.UserService.ConvertToUserResponse(newUser)
			res.AccessResponse = *accessTokenRes
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (u *userImpl) GetMe(context context.Context, id int) (*response.UserResponse, error) {
	user, err := u.Repository.NewUser().FirstRaw(&model.User{ID: uint32(id)})
	if err != nil {
		return nil, err
	}

	return u.UserService.ConvertToUserResponse(user), nil
}

func (u *userImpl) RefreshToken(context context.Context, p *request.RefreshTokenPayload) (*response.LoginResponse, error) {
	var (
		accountDao = u.Repository.NewAccount()
	)
	acc, err := u.Repository.NewAccount().FirstByRefreshToken(p.RefreshToken)
	if err != nil {
		return nil, err
	}

	newToken, newRfToken, err := u.AuthService.GenerateAccessToken(acc.ID, acc.User.Email)
	err = accountDao.InsertAccountRefreshToken(newRfToken)
	if err != nil {
		return nil, err
	}

	userRes := u.UserService.ConvertToUserResponse(acc.User)

	return &response.LoginResponse{
		User:           *userRes,
		AccessResponse: *newToken,
	}, nil
}
