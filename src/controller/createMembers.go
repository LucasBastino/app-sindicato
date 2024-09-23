package controller

import (
	"fmt"
	"math/rand/v2"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

func CreateMembers(c *fiber.Ctx) error {
	var IdEnterprise int
	for i := 1; i < 400; i++ {
		Name := fmt.Sprintf("nombre%d", i)
		// LastName := fmt.Sprintf("apellido%d", i)
		DNI := rand.IntN(30000000) + 20000000
		// Birthday :=
		// Gender := casado soltero y demas
		// Phone := fmt.Sprintf("15%d", rand.IntN(99999999))
		// hacer nombre+apellido+numero@gmail.com
		// Email := fmt.Sprintf("email%d", i)
		// Address := fmt.Sprintf("direccion%d", rand.IntN(9999))
		// PostalCode := strconv.Itoa(rand.IntN(9999))
		// District := fmt.Sprintf("Distrito%d", i)
		// MemberNumber := strconv.Itoa(rand.IntN(99999999))
		IdEnterprise = rand.IntN(100) + 1
		insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI, IdEnterprise) VALUES ('%s','%d', '%d')", Name, DNI, IdEnterprise))
		if err != nil {
			// DBError{"INSERT MEMBER"}.Error(err)
			fmt.Println("error insertando en la DB")
			panic(err)
		}
		insert.Close()
	}
	return c.SendString("hecho")
}
