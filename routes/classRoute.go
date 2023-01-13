package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-praxis-online-class/controllers"
)

func ClassGroup(app *fiber.App) {
	classGroup := app.Group("/class")
	classGroup.Get("/", controllers.ReadClasses)
	classGroup.Get("/:id", controllers.ReadClass)
	classGroup.Post("/", controllers.CreateClass)
	classGroup.Put("/:id", controllers.UpdateClass)
	classGroup.Delete("/:id", controllers.DeleteClass)
}
