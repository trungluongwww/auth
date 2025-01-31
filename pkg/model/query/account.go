package query

import "github.com/trungluongwww/auth/internal/model"

type AccountResult struct {
	model.Account
	AccountRefreshToken []*model.AccountRefreshToken `json:"account_refresh_token"`
}
