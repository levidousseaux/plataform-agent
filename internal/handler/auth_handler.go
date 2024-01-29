package handler

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/levidousseaux/plataform-agent/internal/entity"
	"github.com/shareed2k/goth_fiber"
	"gorm.io/gorm"
)

func MapAuthHandler(app *fiber.App, db *gorm.DB) {
	app.Get("/login/:provider", func(ctx *fiber.Ctx) error {
		url, err := goth_fiber.GetAuthURL(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return ctx.SendString(url)
	})

	app.Get("/auth/callback/:provider", func(ctx *fiber.Ctx) error {
		providerUser, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			log.Println(err)
			ctx.SendStatus(fiber.StatusInternalServerError)
		}

		var user entity.User
		tx := db.First(&user, "email = ? AND provider = ?", providerUser.Email, providerUser.Provider)
		if tx.RowsAffected == 0 {
			user = entity.User{
				Email:    providerUser.Email,
				Name:     providerUser.Name,
				Provider: providerUser.Provider,
			}

			if tx := db.Create(&user); tx.Error != nil {
				return tx.Error
			}
		}

		// Criar um token JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": providerUser.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
		})

		// Assinar o token com a chave secreta
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return err
		}

		// Armazenar o token nos cookies de resposta
		ctx.Cookie(&fiber.Cookie{
			Name:  "jwt",
			Value: tokenString,
		})

		return ctx.Redirect("/")
	})

	app.Get("/logout", func(ctx *fiber.Ctx) error {
		ctx.ClearCookie("jwt")
		return ctx.SendStatus(fiber.StatusOK)
	})
}
