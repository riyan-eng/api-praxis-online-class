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

func ReadUserType(c *fiber.Ctx) error {
	var (
		id       = c.Params("id")
		userType models.UserType
	)

	filter := bson.M{"_id": id}
	err := models.UserTypeCollection().FindOne(c.Context(), filter).Decode(&userType)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data":    "record does not exists",
			"message": "ok",
		})
	} else if err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userType,
		"message": "ok",
	})
}

func UpdateUserType(c *fiber.Ctx) error {
	id := c.Params("id")
	var userType models.UserType

	// validate parsing bodyjson
	if err := c.BodyParser(&userType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// validate body json
	if errorValidate := helpers.ValidateUserType(userType); errorValidate != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errorValidate,
			"message": "fail",
		})
	}

	// access to database
	filter := bson.M{"_id": id}
	update := bson.M{
		"user_type_code": userType.UserTypeCode,
		"user_type_name": userType.UserTypeName,
	}

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := models.UserTypeCollection().FindOneAndUpdate(c.Context(), filter, bson.M{"$set": update}, &opt).Decode(&userType)
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
		"data":    userType,
		"message": "ok",
	})
}

func DeleteUserType(c *fiber.Ctx) error {
	id := c.Params("id")
	var userType models.UserType

	// access to database
	filter := bson.M{"_id": id}
	err := models.UserTypeCollection().FindOneAndDelete(c.Context(), filter).Decode(&userType)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "record does not exists",
			"message": "fail",
		})
	} else if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userType,
		"message": "ok",
	})
}
