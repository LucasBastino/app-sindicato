package controller

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/gofiber/fiber/v2"
)

// func CreateEnterprises(c *fiber.Ctx) error {
// 	var Name string
// 	var Address string
// 	for i := 1; i < 150; i++ {
// 		Name = fmt.Sprintf("empresa%d", i)
// 		Address = fmt.Sprintf("Calle %d", rand.IntN(9999)+1000)
// 		insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s','%s')", Name, Address))
// 		if err != nil {
// 			fmt.Println("error insertando en la DB")
// 			panic(err)
// 		}
// 		insert.Close()
// 	}
// 	return c.SendString("enterprises created")
// }

func CreateEnterprises(c *fiber.Ctx) error {
	type Enterprise struct {
		IdEnterprise int
		Name         string
		Address      string
		CUIT         string
		District     string
		PostalCode   int
		Phone        string
	}
	type JSONData struct {
		EnterprisesName []string
		Streets         []string
	}

	file, err := os.Open("./data/jsonData.json")
	if err != nil {
		fmt.Println("error opening json file")
	}
	decoder := json.NewDecoder(file)
	var jsonData JSONData
	decoder.Decode(&jsonData)

	var e Enterprise
	e.IdEnterprise = rand.IntN(100)
	e.Name = jsonData.EnterprisesName[rand.IntN(len(jsonData.EnterprisesName))]
	m.Address = fmt.Sprintf("%s %d", jsonData.Streets[rand.IntN(len(jsonData.Streets))].Name, rand.IntN(9999))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"enterprise": e})
}
