package controllers

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/api-praxis-online-class/helpers"
	"github.com/riyan-eng/api-praxis-online-class/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	var userType models.UserType
	var class models.Class
	var paymentPeriod models.PaymentPeriod

	// validate require bodyjson
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate bodyjson
	if errValidate := helpers.ValidateUser(user); errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errValidate,
			"message": "fail",
		})
	}

	// find user type
	err := models.UserTypeCollection().FindOne(c.Context(), bson.M{"user_type_code": "participant"}).Decode(&userType)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "user type does not exists",
			"message": "fail",
		})
	} else if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// find class
	err = models.ClassCollection().FindOne(c.Context(), bson.M{"class_code": user.ClassId}).Decode(&class)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "class does not exists",
			"message": "fail",
		})
	} else if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// find payment period
	err = models.PaymentPeriodCollection().FindOne(c.Context(), bson.M{"period_code": user.PaymentPeriodId}).Decode(&paymentPeriod)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "payment code does not exists",
			"message": "fail",
		})
	} else if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// access to database
	user.ID = uuid.New().String()
	user.IsActive = true
	user.UserTypeId = userType.ID
	user.ClassId = class.ID
	user.UniqueCode = uint16(rand.Intn(500))
	user.PaymentPeriodId = paymentPeriod.ID

	result, err := models.UserCollection().InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    result,
		"message": "ok",
	})
}
