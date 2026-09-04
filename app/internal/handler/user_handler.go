package handler

import (
	"user-service/app/internal/domain"
	"user-service/app/internal/helper"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	uc domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

func (h *UserHandler) RegisterUser(c *echo.Context) error {

	return helper.Created(c, "User Created")
}
