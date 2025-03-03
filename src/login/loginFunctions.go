package login

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func getUserAndPassword(c *fiber.Ctx) (string, string) {
	return c.FormValue("user"), c.FormValue("password")
}

func checkUser(user string) bool {
	return !(user != "admin" && user != "guest")
}

func checkPassword(password string) bool {
	return password == "123"
}

func createJwtMapClaims(user string, minutes int) jwt.MapClaims {
	return jwt.MapClaims{
		"role": user,
		"exp":  time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
	}
}

func createToken(claims jwt.MapClaims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func signToken(token *jwt.Token) string {
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("error signing token")
		panic(err)
	}
	return signedToken
}

func createCookie(signedToken string) fiber.Cookie {
	return fiber.Cookie{
		Name:        "Authorization",
		Value:       signedToken,
		Path:        "/",
		Secure:      true,
		HTTPOnly:    true,
		SameSite:    "Strict",
		SessionOnly: true,
	}
}
