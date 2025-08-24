package login

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	pe "github.com/LucasBastino/app-sindicato/src/permissions"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func getUserAndPassword(c *fiber.Ctx) (string, string) {
	return c.FormValue("user"), c.FormValue("password")
}

func checkUser(user string) (string, bool) {
	row := database.DB.QueryRow("SELECT Hash FROM UserTable WHERE Username = ?", user)
	var hash string

	err := row.Scan(&hash)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("username doesn't exist, error:", err)
		} else {
			fmt.Println("error scanning user, error:", err)
		}
		return "", false
	}
	return hash, true
}

func checkHashAndPassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

func createJwtMapClaims(user string, admin bool, p pe.Permissions, hours int) jwt.MapClaims {
	return jwt.MapClaims{
		"user":             user,
		"admin":            admin,
		"writeMember":      p.WriteMember,
		"deleteMember":     p.DeleteMember,
		"writeEnterprise":  p.WriteEnterprise,
		"deleteEnterprise": p.DeleteEnterprise,
		"writeParent":      p.WriteParent,
		"deleteParent":     p.DeleteParent,
		"writePayment":     p.WritePayment,
		"deletePayment":    p.DeletePayment,
		"exp":              time.Now().Add(time.Hour * time.Duration(hours)).Unix(),
	}
}

func createToken(claims jwt.MapClaims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func signToken(token *jwt.Token) (string, *customError.CustomError) {
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return "", &customError.InternalServerError
	}
	return signedToken, nil
}

func createCookie(signedToken string) fiber.Cookie {
	return fiber.Cookie{
		Name:        "Authorization",
		Value:       signedToken,
		Path:        "/",
		HTTPOnly:    true,
		Secure: 	false,
		SameSite:    "Lax",
		SessionOnly: true,
		// para subir a un dominio
		// Secure:   true,
		// SameSite: "None",
	}
}
