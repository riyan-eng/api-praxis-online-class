package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-praxis-online-class/controllers"
)

func UserGroup(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Post("/", controllers.CreateUser)
}
