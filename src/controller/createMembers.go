package controller

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// func CreateMembers(c *fiber.Ctx) error {
// 	var IdEnterprise int
// 	for i := 1; i < 400; i++ {
// 		Name := fmt.Sprintf("nombre%d", i)
// 		// LastName := fmt.Sprintf("apellido%d", i)
// 		DNI := rand.IntN(30000000) + 20000000
// 		// Birthday :=
// 		// Gender := casado soltero y demas
// 		// Phone := fmt.Sprintf("15%d", rand.IntN(99999999))
// 		// hacer nombre+apellido+numero@gmail.com
// 		// Email := fmt.Sprintf("email%d", i)
// 		// Address := fmt.Sprintf("direccion%d", rand.IntN(9999))
// 		// PostalCode := strconv.Itoa(rand.IntN(9999))
// 		// District := fmt.Sprintf("Distrito%d", i)
// 		// MemberNumber := strconv.Itoa(rand.IntN(99999999))
// 		IdEnterprise = rand.IntN(100) + 1
// 		insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI, IdEnterprise) VALUES ('%s','%d', '%d')", Name, DNI, IdEnterprise))
// 		if err != nil {
// 			// DBError{"INSERT MEMBER"}.Error(err)
// 			fmt.Println("error insertando en la DB")
// 			panic(err)
// 		}
// 		insert.Close()
// 	}
// 	return c.SendString("hecho")
// }

func CreateMembers(c *fiber.Ctx) error {
	type Member struct {
		IdMember      int
		Name          string
		LastName      string
		DNI           string
		Birthday      string
		Gender        string
		MaritalStatus string
		Phone         string
		Email         string
		Address       string
		PostalCode    string
		District      string
		MemberNumber  string
		CUIL          string
		IdEnterprise  int
		Category      string
		EntryDate     string
	}

	type JSONData struct {
		FirstNames []string
		LastNames  []string
	}

	member := Member{}
	file, err := os.Open("./data/namesAndLastNames.json")
	if err != nil {
		fmt.Println("error opening file")
		panic(err)
	}
	decoder := json.NewDecoder(file)
	jsonData := JSONData{}
	decoder.Decode(&jsonData)

	member.Name = jsonData.FirstNames[rand.IntN(len(jsonData.FirstNames))]
	member.LastName = jsonData.LastNames[rand.IntN(len(jsonData.LastNames))]
	member.DNI = strconv.Itoa(rand.IntN(3000000) + 2000000)
	year := rand.IntN(114) + 1910
	month := rand.IntN(11) + 1
	var day int
	switch month {
	case 2:
		day = rand.IntN(27) + 1
	case 4, 6, 9, 11:
		day = rand.IntN(29) + 1
	case 1, 3, 5, 7, 8, 10, 12:
		day = rand.IntN(30) + 1
	}
	member.Birthday = fmt.Sprintf("%d/%d/%d", day, month, year)
	fmt.Println(member)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"member": member})
}
