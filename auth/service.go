package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int64) (string, error)
}

type jwtService struct {
	Service
}

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int64) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
