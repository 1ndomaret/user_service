package router

import (
	"user-service/app/internal/handler"

	"github.com/labstack/echo/v5"
)

func Register(e *echo.Echo, userHandler *handler.UserHandler) {
	api := e.Group("/api/v1")
	public := api.Group("")

	public.POST("/users/register", userHandler.RegisterUser)
}
