package usecase_test

import (
	"context"
	"testing"
	"user-service/app/internal/domain"
	"user-service/app/internal/entity"
	"user-service/app/internal/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDonorProfileRepository struct {
	mock.Mock
}

func (m *MockDonorProfileRepository) GetProfile(ctx context.Context, userID uuid.UUID) (*entity.DonorProfile, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.DonorProfile), args.Error(1)
}

func (m *MockDonorProfileRepository) Update(ctx context.Context, req *entity.DonorProfile) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockDonorProfileRepository) Search(ctx context.Context, bloodType, city string) ([]entity.DonorProfile, error) {
	args := m.Called(ctx, bloodType, city)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.DonorProfile), args.Error(1)
}

func (m *MockDonorProfileRepository) GetByID(ctx context.Context, ID uuid.UUID) (*entity.DonorProfile, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.DonorProfile), args.Error(1)
}

func TestGetProfileSuccess(t *testing.T) {
	mockRepo := new(MockDonorProfileRepository)
	uc := usecase.NewDonorProfileUsecase(mockRepo)

	userID := uuid.New()
	expectedProfile := entity.DonorProfile{
		ID:        uuid.New(),
		UserID:    userID,
		BloodType: "O+",
		City:      "Jakarta",
	}

	mockRepo.On("GetProfile", mock.Anything, userID).Return(&expectedProfile, nil)

	res, err := uc.GetProfile(userID)

	assert.NoError(t, err)
	assert.Equal(t, &expectedProfile, res)
	mockRepo.AssertExpectations(t)
}

func TestGetProfileNotFound(t *testing.T) {
	mockRepo := new(MockDonorProfileRepository)
	uc := usecase.NewDonorProfileUsecase(mockRepo)

	userID := uuid.New()

	mockRepo.On("GetProfile", mock.Anything, userID).Return(nil, domain.ErrProfileNotFound)

	res, err := uc.GetProfile(userID)

	assert.Error(t, err)
	assert.Equal(t, domain.ErrProfileNotFound, err)
	assert.Nil(t, res)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfileSuccess(t *testing.T) {
	mockRepo := new(MockDonorProfileRepository)
	uc := usecase.NewDonorProfileUsecase(mockRepo)

	userID := uuid.New()
	initialProfile := entity.DonorProfile{
		ID:          uuid.New(),
		UserID:      userID,
		BloodType:   "O+",
		City:        "Jakarta",
		Latitude:    -6.200000,
		Longitude:   106.816666,
		IsAvailable: false,
	}

	newBloodType := "AB+"
	newCity := "Bekasi"
	newLat := -6.238270
	newLong := 106.975573
	newIsAvailable := true

	req := domain.UpdateProfileRequest{
		BloodType:   &newBloodType,
		City:        &newCity,
		Latitude:    &newLat,
		Longitude:   &newLong,
		IsAvailable: &newIsAvailable,
	}

	mockRepo.On("GetProfile", mock.Anything, userID).Return(&initialProfile, nil)

	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*entity.DonorProfile")).Return(nil)

	res, err := uc.Update(userID, &req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, newBloodType, res.BloodType)
	assert.Equal(t, newCity, res.City)
	assert.Equal(t, newIsAvailable, res.IsAvailable)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfilePartialUpdate(t *testing.T) {
	mockRepo := new(MockDonorProfileRepository)
	uc := usecase.NewDonorProfileUsecase(mockRepo)

	userID := uuid.New()
	initialProfile := entity.DonorProfile{
		ID:          uuid.New(),
		UserID:      userID,
		BloodType:   "O+",
		City:        "Jakarta",
		IsAvailable: false,
	}

	newIsAvailable := true
	req := domain.UpdateProfileRequest{
		IsAvailable: &newIsAvailable,
	}

	mockRepo.On("GetProfile", mock.Anything, userID).Return(&initialProfile, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*entity.DonorProfile")).Return(nil)

	res, err := uc.Update(userID, &req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res.IsAvailable)
	assert.Equal(t, "O+", res.BloodType)
	assert.Equal(t, "Jakarta", res.City)
	mockRepo.AssertExpectations(t)
}

func TestSearchProfileSuccess(t *testing.T) {
	mockRepo := new(MockDonorProfileRepository)
	uc := usecase.NewDonorProfileUsecase(mockRepo)

	mockProfiles := []entity.DonorProfile{
		{ID: uuid.New(), UserID: uuid.New(), BloodType: "AB+", City: "Bekasi", IsAvailable: true},
		{ID: uuid.New(), UserID: uuid.New(), BloodType: "AB+", City: "Bekasi", IsAvailable: true},
	}

	req := domain.SearchProfileRequest{
		BloodType: " AB+ ",
		City:      " Bekasi ",
	}

	mockRepo.On("Search", mock.Anything, "AB+", "Bekasi").Return(mockProfiles, nil)

	res, err := uc.Search(&req)

	assert.NoError(t, err)
	assert.Len(t, res, 2)
	assert.Equal(t, mockProfiles, res)
	mockRepo.AssertExpectations(t)
}

func TestGetByIDSuccess(t *testing.T) {
	mockRepo := new(MockDonorProfileRepository)
	uc := usecase.NewDonorProfileUsecase(mockRepo)

	profileID := uuid.New()
	expectedProfile := entity.DonorProfile{
		ID:        profileID,
		UserID:    uuid.New(),
		BloodType: "A-",
		City:      "Surabaya",
	}

	mockRepo.On("GetByID", mock.Anything, profileID).Return(&expectedProfile, nil)

	res, err := uc.GetByID(profileID)

	assert.NoError(t, err)
	assert.Equal(t, &expectedProfile, res)
	mockRepo.AssertExpectations(t)
}
