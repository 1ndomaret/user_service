package main

import (
	"log"
	"user-service/app/internal/config"
	"user-service/app/internal/entity"
	"user-service/app/internal/handler"
	"user-service/app/internal/repository"
	"user-service/app/internal/router"
	"user-service/app/internal/usecase"

	"github.com/labstack/echo/v5"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	donorProfileRepository := repository.NewDonorProfileRepository(db)
	donorProfileUsecase := usecase.NewDonorProfileUsecase(donorProfileRepository)
	donorProfileHandler := handler.NewDonorProfileHandler(donorProfileUsecase)

	e := echo.New()

	router.Register(e,
		userHandler,
		donorProfileHandler,
	)
	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
