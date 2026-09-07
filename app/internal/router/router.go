package router

import (
	"user-service/app/internal/handler"
	"user-service/app/internal/middleware"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func Register(e *echo.Echo,
	userHandler *handler.UserHandler,
	donorProfileHandler *handler.DonorProfileHandler,
) {
	api := e.Group("/api/v1")
	public := api.Group("")

	public.POST("/users/register", userHandler.Register)
	public.POST("/users/login", userHandler.Login)

	private := api.Group("")
	private.Use(echojwt.WithConfig(middleware.JwtConfig()), middleware.ParseJwtClaims)

	private.GET("/users/donor_profile", donorProfileHandler.GetProfile)
	private.PUT("/users/donor_profile", donorProfileHandler.UpdateProfile)
}
