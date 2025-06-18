package auth

import "github.com/golang-jwt/jwt/v5"

type Authenticator interface {
	GenerateToken(claims jwt.MapClaims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
