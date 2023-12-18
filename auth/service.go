package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID uint64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	Service
}

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

func NewService() *jwtService {
	return &jwtService{}
}

// GenerateToken generates a JWT token for the given user ID.
//
// Parameters:
// - userID: the ID of the user for whom the token is being generated.
//
// Returns:
// - tokenString: the generated JWT token as a string.
// - err: any error encountered during token generation.
func (s *jwtService) GenerateToken(userID uint64) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

// ValidateToken validates a JWT token.
// It takes a token string as input and returns a parsed JWT token and an error if any.
func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t_.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
}
