package query

import "github.com/trungluongwww/auth/internal/model"

type UserResult struct {
	model.User
	Account *model.Account `json:"account"`
}
