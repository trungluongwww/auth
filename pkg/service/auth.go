package service

import (
	"github.com/google/uuid"
	"github.com/trungluongwww/auth/internal/model"
	"github.com/trungluongwww/auth/internal/signer"
	"github.com/trungluongwww/auth/pkg/model/request"
	"github.com/trungluongwww/auth/pkg/model/response"
	"github.com/trungluongwww/auth/util/crypto"
)

type AuthService interface {
	ConvertToAccountModel(password string) (*model.Account, error)
	GenerateAccessToken(accountID uint32, email string) (*response.AccessResponse, *model.AccountRefreshToken, error)
}

type authService struct {
	Signer signer.Signer
}

func NewAuthService(signer signer.Signer) AuthService {
	return &authService{
		Signer: signer,
	}
}

func (authService) ConvertRegisterPayloadToModel(p *request.RegisterPayload, account *model.Account) *model.User {
	return &model.User{
		AccountID: account.ID,
		Email:     p.Email,
		Name:      p.Name,
		IsBanned:  false,
	}
}

func (authService) ConvertToAccountModel(password string) (*model.Account, error) {
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &model.Account{
		Password: hashedPassword,
	}, nil
}

func (authService) CheckAccountLogin(password string, account *model.Account) error {
	err := crypto.CompareHashAndPassword(account.Password, password)
	if err != nil {
		return err
	}
	return nil
}

func (s authService) GenerateAccessToken(accountID uint32, email string) (*response.AccessResponse, *model.AccountRefreshToken, error) {
	ID, _ := uuid.NewUUID()
	accessToken, err := s.Signer.SignUser(ID, int(accountID), email)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, _ := uuid.NewUUID()
	accountRefreshToken := &model.AccountRefreshToken{
		AccountID:     accountID,
		Token:         refreshToken.String(),
		AccessTokenID: ID.String(),
	}

	return &response.AccessResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.String(),
	}, accountRefreshToken, nil
}
