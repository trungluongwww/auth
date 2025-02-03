package service

import (
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
)

type UserService interface {
	ConvertRegisterPayloadToModel(p request.RegisterPayload, account *model.Account, isFirstLogin bool) *model.User
	ConvertToUserResponse(doc *model.User) *response.UserResponse
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (userService) ConvertRegisterPayloadToModel(p request.RegisterPayload, account *model.Account, isFirstLogin bool) *model.User {
	return &model.User{
		AccountID: account.ID,
		Email:     p.Email,
		Name:      p.Name,
		IsBanned:  false,
	}
}

func (userService) ConvertToUserResponse(doc *model.User) *response.UserResponse {
	return &response.UserResponse{
		ID:    int(doc.ID),
		Name:  doc.Name,
		Email: doc.Email,
	}
}
