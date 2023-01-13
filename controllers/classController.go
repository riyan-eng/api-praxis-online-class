package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-praxis-online-class/helpers"
	"github.com/riyan-eng/api-praxis-online-class/models"
)

// type Class struct {
// 	ClassName       string `json:"class_name"`
// 	ClassCode       string `json:"class_code"`
// 	ClassMonthPrice int    `json:"class_month_price"`
// 	IsActive        bool   `json:"is_active"`
// }

func CreateClass(c *fiber.Ctx) error {
	class := new(models.Class)
	if err := c.BodyParser(class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	if errorValidate := helpers.ValidateClass(*class); errorValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errorValidate,
			"message": "fail",
		})
	}

	// body := c.Bind(fiber.Map{class})
	// fmt.Println(body)
	// if err := validate.Struct(body); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"data":    err.Error(),
	// 		"message": "fail",
	// 	})
	// }
	newClass := models.Class{
		ClassName: class.ClassName,
	}
	fmt.Println(newClass)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "create",
		"message": "ok",
	})
}

func ReadClasses(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "list",
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
