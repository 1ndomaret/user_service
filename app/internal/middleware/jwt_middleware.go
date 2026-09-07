package middleware

import (
	"strings"
	"user-service/app/internal/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return helper.Unauthorized(c, "missing authorization header")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return helper.Unauthorized(c, "invalid authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		env, err := godotenv.Read("secret.env")
		if err != nil {
			return helper.InternalServerError(c, "failed to load environment")
		}

		secret := env["JWT_SECRET"]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return helper.Unauthorized(c, "invalid or expired token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return helper.Unauthorized(c, "invalid token claims")
		}

		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])

		return next(c)
	}
}
