package routes

import (
	"echo-app/controllers"
	"echo-app/validations"

	"github.com/labstack/echo/v4"
)

func auth(e *echo.Echo) {
	e.POST("/register", controllers.Register, validations.UserRegister)
	// e.POST("/login", controllers.Login)
	// e.PATCH("/password", controllers.ChangePassword, middlewares.Auth)
	// e.POST("/reset-password", controllers.ResetPassword, middlewares.Auth)
	// e.GET("/:username", controllers.GetUserByUsername)
	// e.PUT("/me", controllers.UpdateUserInfo, middlewares.Auth)
	// e.PATCH("/me/change-avatar", controllers.ChangeAvatar, middlewares.Auth)
}