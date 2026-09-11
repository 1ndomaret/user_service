package repository

import (
	"context"
	"errors"
	"strings"
	"user-service/app/internal/domain"
	"user-service/app/internal/entity"

	"github.com/google/uuid"

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

func (r *donorProfileRepository) GetProfile(ctx context.Context, userID uuid.UUID) (*entity.DonorProfile, error) {
	var profile entity.DonorProfile

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	return &profile, nil
}

func (r *donorProfileRepository) GetByID(ctx context.Context, userID uuid.UUID) (*entity.DonorProfile, error) {
	var profile entity.DonorProfile

	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	return &profile, nil
}

func (r *donorProfileRepository) Update(ctx context.Context, req *entity.DonorProfile) error {
	return r.db.WithContext(ctx).Save(req).Error
}

func (r *donorProfileRepository) Search(ctx context.Context, bloodType, city string) ([]entity.DonorProfile, error) {
	var profiles []entity.DonorProfile

	query := r.db.WithContext(ctx).Model(&entity.DonorProfile{})
	if bloodType != "" {
		query = query.Where("blood_type = UPPER(?)", strings.ToUpper(bloodType))
	}
	if city != "" {
		query = query.Where("city ILIKE ?", "%"+city+"%")
	}

	if err := query.Find(&profiles).Error; err != nil {
		return nil, err
	}

	return profiles, nil
}
