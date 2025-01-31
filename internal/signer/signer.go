package signer

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/trungluongwww/auth/config"
	"strconv"
	"time"
)

const (
	Issuer      = "github.com/trungluongwww/auth"
	ExpiryHours = 24 * 356
)

type Signer interface {
	SignUser(TokenID uuid.UUID, accountID int, email string) (string, error)
}

type signerImpl struct {
	SecretUserKey string
}

func NewSigner(cfg config.Env) Signer {
	return &signerImpl{
		SecretUserKey: cfg.SecretUserJWTToken,
	}
}

func (s *signerImpl) SignUser(TokenID uuid.UUID, accountID int, email string) (string, error) {
	now := time.Now()
	claim := Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Issuer,
			Subject:   strconv.Itoa(accountID),
			Audience:  []string{email},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * ExpiryHours)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        TokenID.String(),
		},
	}

	return s.Sign(claim, s.SecretUserKey)
}

func (s *signerImpl) Sign(claim Claim, key string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(key))
}
