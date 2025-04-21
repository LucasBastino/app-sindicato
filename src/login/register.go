package login

import (
	"database/sql"
	"fmt"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	user := c.FormValue("user")
	password := c.FormValue("password")
	role := c.FormValue("role")

	fmt.Println("user:", user)
	fmt.Println("password:", password)
	fmt.Println("role:", role)

	row := database.DB.QueryRow("SELECT IdUser FROM UserTable WHERE Username = ?", user)
	var idUser int
	err := row.Scan(&idUser)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, error:", err)
		return c.Render("register", fiber.Map{"userError": "Nombre de usuario ya existente"})
	}

	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("error generating hash from password")
		panic(err)
	}

	strHash := string(byteHash)
	fmt.Println(strHash)

	insert, err := database.DB.Query("INSERT INTO UserTable (Username, Hash, Role) VALUES ('?', '?', '?')", user, strHash, role)
	if err != nil {
		fmt.Println("error inserting user in DB")
		panic(err)
	}
	insert.Close()
	return c.Render("register", fiber.Map{})

}
