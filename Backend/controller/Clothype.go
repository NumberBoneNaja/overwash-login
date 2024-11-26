package controller

import (
	"four/config"
	"four/entity"

	"github.com/gofiber/fiber/v2"
)

func CreateClothType(c *fiber.Ctx) error {
	db := config.DB()
	var ClothType entity.ClothType
	if err := c.BodyParser(&ClothType); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := db.Create(&ClothType).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    ClothType,
	})
}