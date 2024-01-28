package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/levidousseaux/plataform-agent/internal/entity"
	"gorm.io/gorm"
)

func MapProjectRoutes(router fiber.Router, db *gorm.DB) {
	group := router.Group("/project")

	group.Get("/", func(c *fiber.Ctx) error {
		var projects []entity.Project
		tx := db.Find(&projects)
		if tx.Error != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(projects)
	})

	group.Post("/", func(c *fiber.Ctx) error {
		project := new(entity.Project)
		if err := c.BodyParser(project); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if tx := db.Create(project); tx.Error != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusCreated)
	})
}
