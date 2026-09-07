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
}

type DonorProfileRepository interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (*entity.DonorProfile, error)
	Update(ctx context.Context, req *UpdateProfileRequest) (*entity.DonorProfile, error)
}

type DonorProfileUsecase interface {
	GetProfile(userID uuid.UUID) (*entity.DonorProfile, error)
	Update(req *UpdateProfileRequest) (*entity.DonorProfile, error)
}
