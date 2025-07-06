package register

import (
	"github.com/trungluongwww/auth/pkg/service"
	"github.com/trungluongwww/auth/pkg/usecase"
)

func (r *Register) NewUsecaseUser() usecase.User {
	return usecase.NewUser(r.Repository, service.NewAuthService(r.Signer), service.NewUserService(), r.Social)
}

func (r *Register) NewUsecasePost() usecase.Post {
	return usecase.NewPost(r.Repository, service.NewPostService(), service.NewUserService())
}
