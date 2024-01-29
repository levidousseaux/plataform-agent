package handler

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func MapWebHandler(router fiber.Router) {
	webPath := os.Getenv("WEB_PATH")
	router.Static("/index.js", webPath+"/index.js")
	router.Static("/index.css", webPath+"/index.css")
	router.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile(webPath + "/index.html")
	})
}
