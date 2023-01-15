package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-praxis-online-class/controllers"
)

func UserTypeGroup(app *fiber.App) {
	UserTypeGroup := app.Group("/user_type")
	// UserTypeGroup.Get("/", controllers.ReadClasses)
	// UserTypeGroup.Get("/:id", controllers.ReadClass)
	UserTypeGroup.Post("/", controllers.CreateUserType)
	// UserTypeGroup.Put("/:id", controllers.UpdateClass)
	// UserTypeGroup.Delete("/:id", controllers.DeleteClass)
}
