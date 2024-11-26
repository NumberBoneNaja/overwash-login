package controller

import (
	"four/config"
	"four/entity"

	"github.com/gofiber/fiber/v2"
)

func CreatePackage(c *fiber.Ctx) error {
	var Package entity.Package
	db := config.DB()
	if err := c.BodyParser(&Package); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := db.Create(&Package).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    Package,
	})

}

