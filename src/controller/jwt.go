package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUser(ctx *fiber.Ctx) error {
	user := ctx.FormValue("user")
	password := ctx.FormValue("password")

	if user != "admin" && user != "guest" {
		return ctx.Render("loginUnsuccesful", fiber.Map{"error": "the user doesn't exist"})
	}
	if password != "123" {
		return ctx.Render("loginUnsuccesful", fiber.Map{"error": "incorrect password"})
	}

	claims := jwt.MapClaims{
		"role": user,
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("error signing token")
		panic(err)
	}

	cookie := fiber.Cookie{
		Name:        "Authorization",
		Value:       signedToken,
		Path:        "/",
		Secure:      true,
		HTTPOnly:    true,
		SameSite:    "Strict",
		SessionOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.Render("loginSuccesful", fiber.Map{})
}

func VerifyToken(c *fiber.Ctx) error {
	tokenStr := c.Cookies("Authorization")
	_, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Render("sessionExpired", fiber.Map{"error": "user unauthorized"})
	}
	return c.Next()
}
