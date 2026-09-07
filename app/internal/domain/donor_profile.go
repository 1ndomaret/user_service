package domain

import (
	"context"
	"user-service/app/internal/entity"
)

type CreateProfileRequest struct {
}

type UpdateProfileRequest struct {
}

type DonorProfileRepository interface {
	Create(ctx context.Context, req *CreateProfileRequest) (*entity.DonorProfile, error)
	GetProfile(ctx context.Context) (*entity.DonorProfile, error)
	Update(ctx context.Context, req *UpdateProfileRequest) (*entity.DonorProfile, error)
}

type DonorProfileUsecase interface {
	GetProfile() (*entity.DonorProfile, error)
	Update(req *UpdateProfileRequest) (*entity.DonorProfile, error)
}
