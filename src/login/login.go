package login

import (
	"os"
	"time"

	"github.com/LucasBastino/app-sindicato/src/config/logger"
	"github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	pe "github.com/LucasBastino/app-sindicato/src/permissions"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUser(c *fiber.Ctx) error {

	user, password := getUserAndPassword(c)
	hash, checkedUser := checkUser(user)
	if !checkedUser {
		return c.Render("login", fiber.Map{"user": user, "password": password, "userError": "Usuario no existente"})
	}

	checkErr := checkHashAndPassword([]byte(hash), []byte(password))
	if checkErr != nil {
		return c.Render("login", fiber.Map{"user": user, "password": password, "passwordError": "Contraseña incorrecta"})
	}
	admin, err := pe.GetAdmin(user)
	if err != nil {
		return errorHandler.HandleError(c, err)
	}
	permissions, err := pe.GetPermissions(user)
	if err != nil {
		return errorHandler.HandleError(c, err)
	}
	claims := createJwtMapClaims(user, admin, permissions, 8)
	token := createToken(claims)
	signedToken, err := signToken(token)
	if err != nil {
		return errorHandler.HandleError(c, err)
	}
	cookie := createCookie(signedToken)
	c.Cookie(&cookie)

	return c.Render("loginSuccesful", fiber.Map{})
}

func RenderInsufficientPermissions(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).Render("insufficientPermissions", fiber.Map{})
}

func RenderExpiredSession(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).Render("expiredSession", fiber.Map{})
}

func VerifyToken(c *fiber.Ctx) error {
	tokenStr := c.Cookies("Authorization")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		logger.Log.Warn("invalid session")
		return c.Render("redirect", fiber.Map{"path": "/expiredSession"})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("claims", claims)
	return c.Next()
}

func LogoutUser(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expired 1 hour ago
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		// para subir a un dominio
		// Secure:   true,
		// SameSite: "None",
	}

	c.Cookie(&cookie)

	return c.Render("login", fiber.Map{})

}


