package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/api-praxis-online-class/helpers"
	"github.com/riyan-eng/api-praxis-online-class/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateClass(c *fiber.Ctx) error {
	class := new(models.Class)
	class.ID = uuid.New().String()
	class.IsActive = true

	// validate require body json
	if err := c.BodyParser(class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if errorValidate := helpers.ValidateClass(*class); errorValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errorValidate,
			"message": "fail",
		})
	}

	// access to database
	result, err := models.ClassCollection().InsertOne(c.Context(), class)
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

func ReadClasses(c *fiber.Ctx) error {
	result, err := models.ClassCollection().Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	defer result.Close(c.Context())

	var classes []models.Class
	for result.Next(c.Context()) {
		var class models.Class
		if err = result.Decode(&class); err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		classes = append(classes, class)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    classes,
		"message": "ok",
	})
}

func ReadClass(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "read",
		"message": "ok",
	})
}

func UpdateClass(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "update",
		"message": "ok",
	})
}

func DeleteClass(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "delete",
		"message": "ok",
	})
}
