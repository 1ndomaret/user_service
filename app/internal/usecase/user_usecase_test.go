package usecase

import (
	"errors"
	"testing"
	"user-service/app/internal/domain"
	"user-service/app/internal/entity"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type fakeUserRepository struct {
	createdUser  *entity.User
	existingUser *entity.User
}

func (f *fakeUserRepository) Create(user *entity.User) error {
	f.createdUser = user
	return nil
}

func (f *fakeUserRepository) GetByEmail(email string) (*entity.User, error) {
	if f.existingUser != nil {
		return f.existingUser, nil
	}

	return nil, errors.New("user not found")
}

func TestRegisterSuccess(t *testing.T) {
	repo := &fakeUserRepository{}

	uc := NewUserUsecase(repo)

	req := domain.RegisterRequest{
		Name:     "Imam",
		Email:    "imam@mail.com",
		Password: "secret123",
		Phone:    "08123456789",
		Role:     "donor",
	}

	user, err := uc.Register(req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Email != req.Email {
		t.Errorf("expected email %s, got %s", req.Email, user.Email)
	}

	if repo.createdUser == nil {
		t.Fatal("expected user to be created")
	}

	if repo.createdUser.Password == req.Password {
		t.Error("password should be hashed")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(repo.createdUser.Password),
		[]byte(req.Password),
	)

	if err != nil {
		t.Error("hashed password does not match original password")
	}
}

func TestRegisterPasswordTooShort(t *testing.T) {
	repo := &fakeUserRepository{}
	uc := NewUserUsecase(repo)

	req := domain.RegisterRequest{
		Name:     "Imam",
		Email:    "imam@mail.com",
		Password: "123",
		Phone:    "08123456789",
		Role:     "donor",
	}

	user, err := uc.Register(req)

	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}

	if user != nil {
		t.Error("expected user to be nil")
	}

	if repo.createdUser != nil {
		t.Error("user should not be created")
	}
}

func TestRegisterInvalidRole(t *testing.T) {
	repo := &fakeUserRepository{}
	uc := NewUserUsecase(repo)

	req := domain.RegisterRequest{
		Name:     "Imam",
		Email:    "imam@mail.com",
		Password: "secret123",
		Phone:    "08123456789",
		Role:     "admin",
	}

	user, err := uc.Register(req)

	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}

	if user != nil {
		t.Error("expected user to be nil")
	}

	if repo.createdUser != nil {
		t.Error("user should not be created")
	}
}

func TestRegisterDuplicateEmail(t *testing.T) {
	repo := &fakeUserRepository{
		existingUser: &entity.User{
			ID:    uuid.New(),
			Name:  "Existing User",
			Email: "imam@mail.com",
		},
	}

	uc := NewUserUsecase(repo)

	req := domain.RegisterRequest{
		Name:     "Imam",
		Email:    "imam@mail.com",
		Password: "secret123",
		Phone:    "08123456789",
		Role:     "donor",
	}

	user, err := uc.Register(req)

	if !errors.Is(err, domain.ErrEmailExists) {
		t.Errorf("expected ErrEmailExists, got %v", err)
	}

	if user != nil {
		t.Error("expected user to be nil")
	}

	if repo.createdUser != nil {
		t.Error("duplicate user should not be created")
	}
}

func TestLoginSuccess(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("secret123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		t.Fatal(err)
	}

	repo := &fakeUserRepository{
		existingUser: &entity.User{
			ID:       uuid.New(),
			Name:     "Imam",
			Email:    "imam@mail.com",
			Password: string(hashedPassword),
			Role:     "donor",
		},
	}

	uc := NewUserUsecase(repo)

	req := domain.LoginRequest{
		Email:    "imam@mail.com",
		Password: "secret123",
	}

	token, err := uc.Login(req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if token == "" {
		t.Error("expected JWT token, got empty string")
	}
}

func TestLoginWrongPassword(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("secret123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		t.Fatal(err)
	}

	repo := &fakeUserRepository{
		existingUser: &entity.User{
			ID:       uuid.New(),
			Email:    "imam@mail.com",
			Password: string(hashedPassword),
			Role:     "donor",
		},
	}

	uc := NewUserUsecase(repo)

	req := domain.LoginRequest{
		Email:    "imam@mail.com",
		Password: "salah123",
	}

	token, err := uc.Login(req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if token != "" {
		t.Error("expected empty token")
	}
}

func TestLoginUserNotFound(t *testing.T) {
	repo := &fakeUserRepository{}

	uc := NewUserUsecase(repo)

	req := domain.LoginRequest{
		Email:    "notfound@mail.com",
		Password: "secret123",
	}

	token, err := uc.Login(req)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if token != "" {
		t.Error("expected empty token")
	}
}
