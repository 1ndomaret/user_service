package usecase

import (
	"context"
	"strings"
	"time"
	"user-service/app/internal/domain"
	"user-service/app/internal/entity"

	"github.com/google/uuid"
)

type donorProfileUsecase struct {
	donorRepo domain.DonorProfileRepository
}

func NewDonorProfileUsecase(donorRepo domain.DonorProfileRepository) domain.DonorProfileUsecase {
	return &donorProfileUsecase{
		donorRepo: donorRepo,
	}
}

var timeOut = 10 * time.Second

func (u *donorProfileUsecase) GetProfile(userID uuid.UUID) (*entity.DonorProfile, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeOut)
	defer cancel()

	return u.donorRepo.GetProfile(ctx, userID)
}

func (u *donorProfileUsecase) Update(userID uuid.UUID, req *domain.UpdateProfileRequest) (*entity.DonorProfile, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeOut)
	defer cancel()

	profile, err := u.donorRepo.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	// nil checks, ignore if nil
	if req.BloodType != nil {
		profile.BloodType = *req.BloodType
	}
	if req.City != nil {
		profile.City = *req.City
	}
	if req.Latitude != nil {
		profile.Latitude = *req.Latitude
	}
	if req.Longitude != nil {
		profile.Longitude = *req.Longitude
	}
	if req.IsAvailable != nil {
		profile.IsAvailable = *req.IsAvailable
	}

	if err := u.donorRepo.Update(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (u *donorProfileUsecase) Search(req *domain.SearchProfileRequest) ([]entity.DonorProfile, error) {
	bloodType := strings.TrimSpace(req.BloodType)
	city := strings.TrimSpace(req.City)

	ctx, cancel := context.WithTimeout(context.TODO(), timeOut)
	defer cancel()

	return u.donorRepo.Search(ctx, bloodType, city)
}

func (u *donorProfileUsecase) GetByID(ID uuid.UUID) (*entity.DonorProfile, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeOut)
	defer cancel()

	return u.donorRepo.GetByID(ctx, ID)
}
