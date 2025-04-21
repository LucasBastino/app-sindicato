package login

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func getUserAndPassword(c *fiber.Ctx) (string, string) {
	return c.FormValue("user"), c.FormValue("password")
}

func checkUser(user string) (string, string, bool) {
	row := database.DB.QueryRow("SELECT Hash, Role FROM UserTable WHERE Username = ?", user)
	var hash string
	var role string
	err := row.Scan(&hash, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("username doesn't exist, error:", err)
		} else {
			fmt.Println("error scanning user, error:", err)
		}
		return "", "", false
	}
	return hash, role, true
}

func checkHashAndPassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

func createJwtMapClaims(user, role string, minutes int) jwt.MapClaims {
	return jwt.MapClaims{
		"user": user,
		"role": role,
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
