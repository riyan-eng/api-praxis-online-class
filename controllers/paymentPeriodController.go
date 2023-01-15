package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-praxis-online-class/models"
)

func CreatePaymentPeriod(c *fiber.Ctx) error {
	var paymentPeriod models.PaymentPeriod
	// validate require bodyjson
	if err := c.BodyParser(&paymentPeriod); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	// validate bodyjson
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}

// func ReadPaymentPeriods(c *fiber.Ctx) error {
// 	var paymentPeriod models.PaymentPeriod
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"data":    1,
// 		"message": "ok",
// 	})
// }

// func ReadPaymentPeriod(c *fiber.Ctx) error {
// 	var paymentPeriod models.PaymentPeriod
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"data":    1,
// 		"message": "ok",
// 	})
// }

// func UpdatePaymentPeriod(c *fiber.Ctx) error {
// 	var paymentPeriod models.PaymentPeriod
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"data":    1,
// 		"message": "ok",
// 	})
// }

// func DeletePaymentPeriod(c *fiber.Ctx) error {
// 	var paymentPeriod models.PaymentPeriod
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"data":    1,
// 		"message": "ok",
// 	})
// }
