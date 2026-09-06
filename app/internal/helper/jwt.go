package helper

import (
	"errors"
	"os"
	"time"
	"user-service/app/internal/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(user *entity.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		env, err := godotenv.Read("secret.env")
		if err != nil {
			return "", err
		}

		secret = env["JWT_SECRET"]
	}

	if secret == "" {
		return "", errors.New("JWT_SECRET is not configured")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
