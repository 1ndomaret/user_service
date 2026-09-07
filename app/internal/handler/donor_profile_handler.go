package handler

import (
	"user-service/app/internal/domain"
	"user-service/app/internal/helper"

	"github.com/labstack/echo/v5"
)

type DonorProfileHandler struct {
	donorProfileUse domain.DonorProfileUsecase
}

func NewDonorProfileHandler(donorProfileUse domain.DonorProfileUsecase) *DonorProfileHandler {
	return &DonorProfileHandler{
		donorProfileUse: donorProfileUse,
	}
}

func (h *DonorProfileHandler) GetProfile(c *echo.Context) error {
	return helper.Success(c, 200, "Success", "Profile")
}

func (h *DonorProfileHandler) UpdateProfile(c *echo.Context) error {
	return helper.Success(c, 200, "Updated", "Profile")
}
