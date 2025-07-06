package usecase

import (
	"context"

	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
	"github.com/trungluongwww/auth/pkg/repository"
	"github.com/trungluongwww/auth/pkg/service"
	"github.com/trungluongwww/auth/third_party/social"
)

// User interface contains all user-related operations
type User interface {
	// Find operations
	GetMe(context context.Context, id int) (*response.UserResponse, error)

	// Upsert operations
	Register(context context.Context, p request.RegisterPayload) error
	Login(context context.Context, p request.LoginPayload) (*response.LoginResponse, error)
	RefreshToken(context context.Context, p *request.RefreshTokenPayload) (*response.LoginResponse, error)
	LoginWithFacebook(context context.Context, p request.FacebookLoginPayload) (*response.LoginResponse, error)
}

type userImpl struct {
	Repository  repository.Repository
	AuthService service.AuthService
	UserService service.UserService
	Social      social.Social
}

// NewUser creates a new User use case
func NewUser(repository repository.Repository, authService service.AuthService, userService service.UserService,
	social social.Social) User {
	return &userImpl{
		Repository:  repository,
		AuthService: authService,
		UserService: userService,
		Social:      social,
	}
}
