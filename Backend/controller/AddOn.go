package controller

import (
	"four/config"
	"four/entity"

	"github.com/gofiber/fiber/v2"
)

func CreateAddOn(c *fiber.Ctx) error {
	db := config.DB()
	var AddOn entity.AddOn
	if err := c.BodyParser(&AddOn); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := db.Create(&AddOn).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    AddOn,
	})
}