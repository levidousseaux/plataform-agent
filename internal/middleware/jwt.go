package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func VerifyAuth(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	if cookie == "" {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println("Parsing token error:", err)
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	// Token válido, você pode acessar as reivindicações agora
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	return ctx.Next()
}
