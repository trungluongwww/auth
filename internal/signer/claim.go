package signer

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	jwt.RegisteredClaims
}
