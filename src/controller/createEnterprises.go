package controller

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/models"
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
	type StreetData struct {
		Name string
	}
	type JSONData struct {
		EnterprisesName []string
		Streets         []StreetData
	}

	file, err := os.Open("./data/jsonData.json")
	if err != nil {
		fmt.Println("error opening json file")
	}
	decoder := json.NewDecoder(file)
	var jsonData JSONData
	decoder.Decode(&jsonData)

	var e models.Enterprise
	for i := 0; i < 50; i++ {
		e.Name = jsonData.EnterprisesName[rand.IntN(len(jsonData.EnterprisesName))]
		e.EnterpriseNumber = strconv.Itoa(rand.IntN(8000) + 1000)
		e.Address = fmt.Sprintf("%s %d", jsonData.Streets[rand.IntN(len(jsonData.Streets))].Name, rand.IntN(9999))
		e.CUIT = fmt.Sprintf("%d-%d-%d", rand.IntN(9)+20, rand.IntN(8999999)+1000000, rand.IntN(8)+1)
		e.District = jsonData.Streets[rand.IntN(len(jsonData.Streets))].Name
		e.PostalCode = strconv.Itoa(rand.IntN(8000) + 1000)
		e.Phone = fmt.Sprintf("156%d", rand.IntN(9999999))

		insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, EnterpriseNumber, Address, CUIT, District, PostalCode, Phone) VALUES ('%s','%s','%s','%s','%s','%s', '%s')", e.Name, e.EnterpriseNumber, e.Address, e.CUIT, e.District, e.PostalCode, e.Phone))
		if err != nil {
			fmt.Println("error inserting enterprise")
			panic(err)
		}
		insert.Close()
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "hecho"})
}
