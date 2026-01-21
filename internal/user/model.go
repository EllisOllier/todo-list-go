package user

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           int     `json:"id"`
	Username     *string `json:"username"`
	PasswordHash *string `json:"-"` // stops the return of password_hash
}

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(givenUserRepository *UserRepository) *UserService {
	return &UserService{
		userRepository: givenUserRepository,
	}
}

func (s *UserService) GenerateToken(user_id int) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	type Claims struct {
		UserId int `json:"user_id"`
		jwt.RegisteredClaims
	}

	claims := Claims{
		user_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Subject:   string(rune(user_id)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)

	return ss, err
}
