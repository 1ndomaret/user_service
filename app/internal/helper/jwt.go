package helper

import (
	"time"
	"user-service/app/internal/config"
	"user-service/app/internal/entity"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User) (string, error) {
	claims := &JwtCustomClaims{
		user.ID,
		user.Email,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := config.JWTSecret()
	if err != nil {
		return "", err
	}

	return token.SignedString(secret)
}
