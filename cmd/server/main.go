package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/levidousseaux/plataform-agent/internal/entity"
	"github.com/levidousseaux/plataform-agent/internal/routes"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	dbUrl := "host=localhost user=postgres password=Teste123@ dbname=platform port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Application{}, &entity.Project{})
	if err != nil {
		log.Fatal(err)
	}

	routes.MapApplicationRoutes(api, db)
	routes.MapProjectRoutes(api, db)

	app.Static("/index.js", "./web/public/index.js")
	app.Static("/index.css", "./web/public/index.css")
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/public/index.html")
	})

	log.Fatal(app.Listen("localhost:3000"))
}
