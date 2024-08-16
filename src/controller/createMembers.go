package controller

import (
	"fmt"
	"math/rand/v2"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

func CreateMembers(c *fiber.Ctx) error {
	var Name string
	var DNI int
	for i := 100; i < 105; i++ {
		Name = fmt.Sprintf("ejemplo%d", i)
		DNI = rand.IntN(30000000) + 20000000
		insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI) VALUES ('%s','%d')", Name, DNI))
		if err != nil {
			// DBError{"INSERT MEMBER"}.Error(err)
			fmt.Println("error insertando en la DB")
			panic(err)
		}
		insert.Close()
	}
	return c.SendString("hecho")
}
