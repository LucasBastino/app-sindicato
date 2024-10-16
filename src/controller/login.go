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

func VerifyAdmin(c *fiber.Ctx) error {
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if role == "admin" {
		return c.Next()
	} else {
		return c.Render("sessionExpired", fiber.Map{"error": "user is not admin"})
	}
}

func VerifyAdminOrGuest(c *fiber.Ctx) error {
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if role == "admin" || role == "guest" {
		return c.Next()
	}
	return c.Render("sessionExpired", fiber.Map{"error": "user is not admin or guest"})
}

func VerifyToken(c *fiber.Ctx) error {
	tokenStr := c.Cookies("Authorization")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Render("sessionExpired", fiber.Map{"error": "session expired"})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("claims", claims)
	return c.Next()
}
