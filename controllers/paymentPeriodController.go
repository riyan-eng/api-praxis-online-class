package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/api-praxis-online-class/helpers"
	"github.com/riyan-eng/api-praxis-online-class/models"
)

func CreatePaymentPeriod(c *fiber.Ctx) error {
	var paymentPeriod models.PaymentPeriod
	paymentPeriod.ID = uuid.New().String()
	paymentPeriod.IsActive = true
	// validate require bodyjson
	if err := c.BodyParser(&paymentPeriod); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	// validate bodyjson
	errValidation := helpers.ValidatePaymentPeriod(paymentPeriod)
	if errValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errValidation,
			"message": "fail",
		})
	}
	// access to database
	result, err := models.PaymentPeriodCollection().InsertOne(c.Context(), paymentPeriod)
	if err != nil {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    result,
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
