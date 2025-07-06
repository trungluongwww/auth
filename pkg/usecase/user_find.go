package usecase

import (
	"context"

	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/response"
)

// GetMe implementation for find operations
func (u *userImpl) GetMe(context context.Context, id int) (*response.UserResponse, error) {
	user, err := u.Repository.NewUser().FirstRaw(&model.User{ID: uint32(id)})
	if err != nil {
		return nil, err
	}

	return u.UserService.ConvertToUserResponse(user), nil
}
