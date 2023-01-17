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

type User struct {
	ID              string               `bson:"_id" json:"id"`
	UserTypeId      models.UserType      `bson:"user_type" json:"user_type"`
	ClassId         models.Class         `bson:"class" json:"class"`
	Name            string               `bson:"name" json:"name"`
	Birthday        string               `bson:"birthday" json:"birthday"`
	Phone           string               `bson:"phone" json:"phone"`
	Email           string               `bson:"email" json:"email"`
	Address         string               `bson:"address" json:"address"`
	Education       string               `bson:"education" json:"education"`
	Reference       string               `bson:"reference" json:"reference"`
	PaymentPeriodId models.PaymentPeriod `bson:"payment_period" json:"payment_period"`
	UniqueCode      uint16               `bson:"unique_code" json:"unique_code"`
	Batch           string               `bson:"batch" json:"batch"`
	IsActive        bool                 `bson:"is_active" json:"is_active"`
}

func ReadUsers(c *fiber.Ctx) error {
	var user User
	var class models.Class
	var userType models.UserType
	var paymentPeriod models.PaymentPeriod
	var users []User

	// access to database
	result, err := models.UserCollection().Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}
	defer result.Close(c.Context())

	for result.Next(c.Context()) {
		var userTemp models.User
		if err := result.Decode(&userTemp); err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		err := models.ClassCollection().FindOne(c.Context(), bson.M{"_id": userTemp.ClassId}).Decode(&class)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		err = models.UserTypeCollection().FindOne(c.Context(), bson.M{"_id": userTemp.UserTypeId}).Decode(&userType)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		err = models.PaymentPeriodCollection().FindOne(c.Context(), bson.M{"_id": userTemp.PaymentPeriodId}).Decode(&paymentPeriod)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		user.ID = userTemp.ID
		user.ClassId = class
		user.UserTypeId = userType
		user.PaymentPeriodId = paymentPeriod
		user.Address = userTemp.Address
		user.Batch = userTemp.Batch
		user.Birthday = userTemp.Birthday
		user.Name = userTemp.Name
		user.Phone = userTemp.Phone
		user.Email = userTemp.Email
		user.Education = userTemp.Education
		user.Reference = userTemp.Reference
		user.UniqueCode = userTemp.UniqueCode
		user.IsActive = userTemp.IsActive
		users = append(users, user)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    users,
		"message": "ok",
	})
}
