package usecase

import (
	"strings"
	"user-service/app/internal/domain"
	"user-service/app/internal/entity"
	"user-service/app/internal/helper"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Register(req domain.RegisterRequest) (*entity.User, error) {
	if strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.Email) == "" ||
		strings.TrimSpace(req.Phone) == "" ||
		strings.TrimSpace(req.Password) == "" {
		return nil, domain.ErrInvalidInput
	}

	if len(req.Password) < 8 {
		return nil, domain.ErrInvalidInput
	}

	if req.Role != "donor" && req.Role != "requester" {
		return nil, domain.ErrInvalidInput
	}

	existingUser, err := u.userRepository.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Phone:    req.Phone,
		Role:     req.Role,
	}

	if err := u.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Login(req domain.LoginRequest) (string, error) {
	user, err := u.userRepository.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return "", err
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
