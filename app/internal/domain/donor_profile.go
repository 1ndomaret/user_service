package domain

import (
	"context"
	"errors"
	"user-service/app/internal/entity"

	"github.com/google/uuid"
)

var (
	ErrProfileNotFound = errors.New("profile not found")
)

type UpdateProfileRequest struct {
	BloodType   *string  `json:"blood_type,omitempty"`
	City        *string  `json:"city,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	IsAvailable *bool    `json:"is_available,omitempty"`
}

type SearchProfileRequest struct {
	BloodType string `query:"blood_type"`
	City      string `query:"city"`
}

type DonorProfileRepository interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (*entity.DonorProfile, error)
	Update(ctx context.Context, req *entity.DonorProfile) error
	Search(ctx context.Context, bloodType, city string) ([]entity.DonorProfile, error)
}

type DonorProfileUsecase interface {
	GetProfile(userID uuid.UUID) (*entity.DonorProfile, error)
	Update(userID uuid.UUID, req *UpdateProfileRequest) (*entity.DonorProfile, error)
	Search(req *SearchProfileRequest) ([]entity.DonorProfile, error)
}
