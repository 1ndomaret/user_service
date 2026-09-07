package domain

import (
	"errors"
	"user-service/app/internal/entity"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrEmailExists  = errors.New("email already registered")
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
}

type UserUsecase interface {
	Register(req RegisterRequest) (*entity.User, error)
	Login(req LoginRequest) (string, error)
}
