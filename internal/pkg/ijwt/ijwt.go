package ijwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTHandler interface {
	SetJWTToken(cp ClaimParams) (string, error)
	ParseToken(tokenStr string) error
}

type JWTHandlerImpl struct {
	Secret []byte
}

func NewJWTHandler(secret string) JWTHandler {
	return &JWTHandlerImpl{
		Secret: []byte(secret),
	}
}

func (r *JWTHandlerImpl) SetJWTToken(cp ClaimParams) (string, error) {
	uc := UserClaims{
		Username: cp.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &uc)
	tokenStr, err := token.SignedString(r.Secret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (r *JWTHandlerImpl) ParseToken(tokenStr string) error {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(r.Secret), nil
	})
	if err != nil {
		return err
	}
	if token.Valid {
		return nil
	}
	return fmt.Errorf("invalid token")
}

type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}
type ClaimParams struct {
	Username string `json:"username"`
}
