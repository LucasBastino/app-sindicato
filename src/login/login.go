package login

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUser(c *fiber.Ctx) error {

	user, password := getUserAndPassword(c)

	hash, role, checkedUser := checkUser(user)

	if !checkedUser {
		return c.Render("login", fiber.Map{"user": user, "password": password, "userError": "Usuario no existente"})
	}

	err := checkHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return c.Render("login", fiber.Map{"user": user, "password": password, "passwordError": "Contraseña incorrecta"})
	}

	claims := createJwtMapClaims(user, role, 90)
	token := createToken(claims)
	signedToken := signToken(token)
	cookie := createCookie(signedToken)
	c.Cookie(&cookie)

	return c.Render("loginSuccesful", fiber.Map{})
}

func RenderInsufficientPermissions(c *fiber.Ctx) error {
	return c.Render("insufficientPermissions", fiber.Map{})
}

func RenderExpiredSession(c *fiber.Ctx) error {
	return c.Render("expiredSession", fiber.Map{})
}

func VerifyToken(c *fiber.Ctx) error {
	tokenStr := c.Cookies("Authorization")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Render("redirect", fiber.Map{"path": "/expiredSession"})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("claims", claims)
	return c.Next()
}

func VerifyAdmin(c *fiber.Ctx) error {
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if role == "admin" {
		return c.Next()
	} else {
		return c.Redirect("/insufficientPermissions")
	}
}

func VerifyAdminOrUser(c *fiber.Ctx) error {
	role := c.Locals("claims").(jwt.MapClaims)["role"]
	if role == "admin" || role == "user" {
		return c.Next()
	}
	return c.Redirect("/insufficientPermissions")
}

func LogoutUser(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expired 1 hour ago
		HTTPOnly: true,
		Secure:   true,
	}

	c.Cookie(&cookie)

	return c.Render("login", fiber.Map{})

}
