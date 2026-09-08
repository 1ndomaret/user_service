package handler

import (
	"strings"
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
	userID, ok := c.Get("user_id").(uuid.UUID)
	if !ok {
		return helper.Unauthorized(c, "invalid or missing token")
	}

	var req domain.UpdateProfileRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	profile, err := h.donorUsecase.Update(userID, &req)
	if err != nil {
		return helper.BadRequest(c, err.Error())
	}

	return helper.Success(c, 200, "Updated", profile)
}

func (h *DonorProfileHandler) Search(c *echo.Context) error {
	req := domain.SearchProfileRequest{
		BloodType: c.QueryParam("blood_type"),
		City:      c.QueryParam("city"),
	}

	if strings.HasSuffix(req.BloodType, " ") {
		req.BloodType = strings.TrimSpace(req.BloodType) + "+"
	}

	profiles, err := h.donorUsecase.Search(&req)
	if err != nil {
		return helper.InternalServerError(c, err.Error())
	}

	return helper.Success(c, 200, "Success", profiles)
}
