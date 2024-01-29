package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/levidousseaux/plataform-agent/internal/entity"
	"gorm.io/gorm"
)

func MapApplicationHandler(router fiber.Router, db *gorm.DB) {
	group := router.Group("/application")
	group.Get("/", func(c *fiber.Ctx) error {
		var applications []entity.Application
		tx := db.Find(&applications)
		if tx.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(applications)
	})

	group.Post("/", func(c *fiber.Ctx) error {
		application := new(entity.Application)
		if err := c.BodyParser(application); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if tx := db.Create(application); tx.Error != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusCreated)
	})
}
