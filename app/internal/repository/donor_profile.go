package repository

import (
	"context"
	"user-service/app/internal/domain"
	"user-service/app/internal/entity"

	"gorm.io/gorm"
)

type donorProfileRepository struct {
	db *gorm.DB
}

func NewDonorProfileRepository(db *gorm.DB) domain.DonorProfileRepository {
	return &donorProfileRepository{
		db: db,
	}
}

func (r *donorProfileRepository) Create(ctx context.Context, req *domain.CreateProfileRequest) (*entity.DonorProfile, error) {
	return nil, nil
}

func (r *donorProfileRepository) GetProfile(ctx context.Context) (*entity.DonorProfile, error) {
	return nil, nil
}

func (r *donorProfileRepository) Update(ctx context.Context, req *domain.UpdateProfileRequest) (*entity.DonorProfile, error) {
	return nil, nil
}
