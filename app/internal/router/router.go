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

	// public.GET("/users/donor-profile/:id", donorProfileHandler.GetByID)
	// public.GET("/users/donor-profile/search", donorProfileHandler.Search)

	private := api.Group("")
	private.Use(echojwt.WithConfig(middleware.JwtConfig()), middleware.ParseJwtClaims)

	private.GET("/users/donor-profile", donorProfileHandler.GetProfile)
	private.PUT("/users/donor-profile", donorProfileHandler.UpdateProfile)

	serviceAuth := api.Group("")
	serviceAuth.Use(middleware.StaticTokenAuth)

	serviceAuth.GET("/users/donor-profile/:id", donorProfileHandler.GetByID)
	serviceAuth.GET("/users/donor-profile/search", donorProfileHandler.Search)

}
