package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/api-praxis-online-class/helpers"
	"github.com/riyan-eng/api-praxis-online-class/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func ReadPaymentPeriods(c *fiber.Ctx) error {
	var paymentPeriod models.PaymentPeriod
	var paymentPeriods []models.PaymentPeriod

	// access to database
	result, err := models.PaymentPeriodCollection().Find(c.Context(), bson.M{})
	if err != nil {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	for result.Next(c.Context()) {
		if err := result.Decode(&paymentPeriod); err != nil {
			c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		paymentPeriods = append(paymentPeriods, paymentPeriod)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    paymentPeriods,
		"message": "ok",
	})
}

func ReadPaymentPeriod(c *fiber.Ctx) error {
	var paymentPeriod models.PaymentPeriod
	var id = c.Params("id")

	// access to database
	filter := bson.M{"_id": id}
	err := models.PaymentPeriodCollection().FindOne(c.Context(), filter).Decode(&paymentPeriod)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "record does not exists",
			"message": "fail",
		})
	} else if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    paymentPeriod,
		"message": "ok",
	})
}

func UpdatePaymentPeriod(c *fiber.Ctx) error {
	var paymentPeriod models.PaymentPeriod
	var id = c.Params("id")

	// validate require bodyjson
	if err := c.BodyParser(&paymentPeriod); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate bodyjson
	if errValidation := helpers.ValidatePaymentPeriod(paymentPeriod); errValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errValidation,
			"message": "fail",
		})
	}

	// access to database
	filter := bson.M{"_id": id}
	update := bson.M{
		"period_code": paymentPeriod.PeriodCode,
		"period_name": paymentPeriod.PeriodName,
		"discount":    paymentPeriod.Discount,
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := models.PaymentPeriodCollection().FindOneAndUpdate(c.Context(), filter, bson.M{"$set": update}, &opt).Decode(&paymentPeriod)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "record does not exists",
			"message": "fail",
		})
	} else if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    paymentPeriod,
		"message": "ok",
	})
}

// func DeletePaymentPeriod(c *fiber.Ctx) error {
// 	var paymentPeriod models.PaymentPeriod
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"data":    1,
// 		"message": "ok",
// 	})
// }
