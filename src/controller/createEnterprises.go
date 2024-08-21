package controller

import (
	"fmt"
	"math/rand/v2"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

func CreateEnterprises(c *fiber.Ctx) error {
	var Name string
	var Address string
	for i := 1; i < 150; i++ {
		Name = fmt.Sprintf("empresa%d", i)
		Address = fmt.Sprintf("Calle %d", rand.IntN(9999)+1000)
		insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s','%s')", Name, Address))
		if err != nil {
			fmt.Println("error insertando en la DB")
			panic(err)
		}
		insert.Close()
	}
	return c.SendString("enterprises created")
}
