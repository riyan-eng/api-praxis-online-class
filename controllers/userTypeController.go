package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/api-praxis-online-class/helpers"
	"github.com/riyan-eng/api-praxis-online-class/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserType(c *fiber.Ctx) error {
	var userType models.UserType
	userType.ID = uuid.New().String()
	userType.IsActive = true

	// validate require body json
	if err := c.BodyParser(&userType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if errorValidate := helpers.ValidateUserType(userType); errorValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errorValidate,
			"message": "fail",
		})
	}

	// access to database
	result, err := models.UserTypeCollection().InsertOne(c.Context(), userType)
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

func ReadUserTypes(c *fiber.Ctx) error {
	result, err := models.UserTypeCollection().Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	defer result.Close(c.Context())

	var userTypes []models.UserType
	for result.Next(c.Context()) {
		var userType models.UserType
		if err = result.Decode(&userType); err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		userTypes = append(userTypes, userType)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userTypes,
		"message": "ok",
	})
}
