package handler

import (
	"errors"
	"user-service/app/internal/domain"
	"user-service/app/internal/helper"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(c *echo.Context) error {
	var req domain.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	user, err := h.userUsecase.Register(req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidInput) {
			return helper.BadRequest(c, "name, email, phone, password, and valid role are required")
		}

		if errors.Is(err, domain.ErrEmailExists) {
			return helper.BadRequest(c, "email already registered")
		}

		return helper.InternalServerError(c, err.Error())
	}

	return helper.Created(c, user)
}

func (h *UserHandler) Login(c *echo.Context) error {
	var req domain.LoginRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	token, err := h.userUsecase.Login(req)
	if err != nil {
		return helper.Unauthorized(c, "invalid email or password")
	}

	return helper.Success(c, 200, "Login Success", map[string]interface{}{
		"token": token,
	})
}
