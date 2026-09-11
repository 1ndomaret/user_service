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

// @Summary   	 Get a user's donor profile
// @Tags         DonorProfile
// @Accept       json
// @Produce      json
// @Success      200      {object}  helper.SwaggoResponse{data=entity.DonorProfile}
// @Failure      400      {object}  helper.ErrorResponse
// @Failure      422      {object}  helper.ErrorResponse
// @Failure      500      {object}  helper.ErrorResponse
// @Security     BearerAuth
// @Router       /users/donor-profile [get]
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

// @Summary   	 Update a user's donor profile
// @Tags         DonorProfile
// @Accept       json
// @Produce      json
// @Param        request  body      domain.UpdateProfileRequest  true  "User Donor Profile Details"
// @Success      200      {object}  helper.SwaggoResponse{data=entity.DonorProfile}
// @Failure      400      {object}  helper.ErrorResponse
// @Failure      422      {object}  helper.ErrorResponse
// @Failure      500      {object}  helper.ErrorResponse
// @Security     BearerAuth
// @Router       /users/donor-profile [put]
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

// @Summary   	 Get donor profiles based on city and blood type
// @Tags         Service
// @Accept       json
// @Produce      json
// @Param 			 blood_type query 		string	 false	 "Blood Type"
// @Param 			 city       query 		string 	 false 	 "City name"
// @Success      200      	{object}  helper.SwaggoResponse{data=[]entity.DonorProfile}
// @Failure      400      	{object}  helper.ErrorResponse
// @Failure      422      	{object}  helper.ErrorResponse
// @Failure      500      	{object}  helper.ErrorResponse
// @Security     BearerAuth
// @Router       /users/donor-profile/search [get]
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

// @Summary   	 Get donor profile by ID
// @Tags         Service
// @Accept       json
// @Produce      json
// @Param        id       path      string true "Donor Profile ID" format(uuid)
// @Success      200      {object}  helper.SwaggoResponse{data=entity.DonorProfile}
// @Failure      400      {object}  helper.ErrorResponse
// @Failure      422      {object}  helper.ErrorResponse
// @Failure      500      {object}  helper.ErrorResponse
// @Security     BearerAuth
// @Router       /users/donor-profile/{id} [get]
func (h *DonorProfileHandler) GetByID(c *echo.Context) error {
	donorID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return helper.BadRequest(c, "invalid blood request id")
	}

	profile, err := h.donorUsecase.GetByID(donorID)
	if err != nil {
		return helper.InternalServerError(c, err.Error())
	}

	return helper.Success(c, 200, "Success", profile)
}
