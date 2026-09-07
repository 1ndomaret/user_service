package usecase

import (
	"context"
	"time"
	"user-service/app/internal/domain"
	"user-service/app/internal/entity"
)

type donorProfileUsecase struct {
	donorProfileRepo domain.DonorProfileRepository
}

func NewDonorProfileUsecase(donorProfileRepo domain.DonorProfileRepository) domain.DonorProfileUsecase {
	return &donorProfileUsecase{
		donorProfileRepo: donorProfileRepo,
	}
}

var timeOut = 10 * time.Second

func (u *donorProfileUsecase) GetProfile() (*entity.DonorProfile, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeOut)
	defer cancel()

	return u.donorProfileRepo.GetProfile(ctx)
}

func (u *donorProfileUsecase) Update(req *domain.UpdateProfileRequest) (*entity.DonorProfile, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeOut)
	defer cancel()

	return u.donorProfileRepo.Update(ctx, req)
}
