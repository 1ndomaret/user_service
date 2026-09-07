package handler

import (
	"user-service/app/internal/domain"
	"user-service/app/internal/helper"

	"github.com/google/uuid"

	"github.com/labstack/echo/v5"
)

type DonorProfileHandler struct {
	donorUsecase domain.DonorProfileUsecase
}

func NewDonorProfileHandler(donorUsecase domain.DonorProfileUsecase) *DonorProfileHandler {
	return &DonorProfileHandler{
		donorUsecase: donorUsecase,
	}
}

func (h *DonorProfileHandler) GetProfile(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uuid.UUID)
	if !ok {
		return helper.Unauthorized(c, "invalid or missing token")
	}

	profile, err := h.donorUsecase.GetProfile(userID)
	if err != nil {
		return helper.InternalServerError(c, err.Error())
	}
	return helper.Success(c, 200, "Success", profile)
}

func (h *DonorProfileHandler) UpdateProfile(c *echo.Context) error {
	return helper.Success(c, 200, "Updated", "Profile")
}
