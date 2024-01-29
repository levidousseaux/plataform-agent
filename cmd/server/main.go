package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/levidousseaux/plataform-agent/internal/entity"
	"github.com/levidousseaux/plataform-agent/internal/handler"
	"github.com/levidousseaux/plataform-agent/internal/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file loaded, using environment variables")
	}

	db := initializeDatabase()
	app := fiber.New()
	initializeAuth()

	api := app.Group("/api", middleware.VerifyAuth)
	handler.MapAuthHandler(app, db)
	handler.MapApplicationHandler(api, db)
	handler.MapProjectHandler(api, db)
	handler.MapWebHandler(app)

	log.Fatal(app.Listen("localhost:3000"))
}

func initializeAuth() {
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("GOOGLE_CALLBACK_URL"),
			"email",
			"profile",
		),
	)
}

func initializeDatabase() *gorm.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&entity.Application{},
		&entity.Project{},
		&entity.User{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
