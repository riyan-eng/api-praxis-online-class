package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-praxis-online-class/controllers"
)

func PaymentPeriodGroup(app *fiber.App) {
	paymentPeriodGroup := app.Group("/payment_period")
	paymentPeriodGroup.Post("/", controllers.CreatePaymentPeriod)
	paymentPeriodGroup.Get("/", controllers.ReadPaymentPeriods)
	paymentPeriodGroup.Get("/:id", controllers.ReadPaymentPeriod)
	paymentPeriodGroup.Put("/:id", controllers.UpdatePaymentPeriod)
	paymentPeriodGroup.Delete(":id", controllers.DeletePaymentPeriod)
}
