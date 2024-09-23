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

	type StreetData struct {
		Code string
		Name string
	}
	type JSONData struct {
		MaleFirstNames   []string
		FemaleFirstNames []string
		LastNames        []string
		MaritalStatus    []string
		Genders          []string
		Streets          []StreetData
		Categories       []string
	}

	m := Member{}
	file, err := os.Open("./data/jsonData.json")
	if err != nil {
		fmt.Println("error opening file")
		panic(err)
	}
	decoder := json.NewDecoder(file)
	jsonData := JSONData{}
	decoder.Decode(&jsonData)

	m.IdMember = rand.IntN(1000)
	m.Gender = jsonData.Genders[rand.IntN(len(jsonData.Genders))]
	if m.Gender == "Hombre" {
		m.Name = jsonData.MaleFirstNames[rand.IntN(len(jsonData.MaleFirstNames))]
	} else if m.Gender == "Mujer" {
		m.Name = jsonData.FemaleFirstNames[rand.IntN(len(jsonData.FemaleFirstNames))]
	} else if m.Gender == "Otro" {
		random := rand.IntN(1)
		if random == 0 {
			m.Name = jsonData.MaleFirstNames[rand.IntN(len(jsonData.MaleFirstNames))]
		} else {
			m.Name = jsonData.FemaleFirstNames[rand.IntN(len(jsonData.FemaleFirstNames))]
		}
	}
	m.LastName = jsonData.LastNames[rand.IntN(len(jsonData.LastNames))]
	m.DNI = strconv.Itoa(rand.IntN(3000000) + 2000000)
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
	// fijarse bien desp lo del formato fecha
	m.Birthday = fmt.Sprintf("%d/%d/%d", day, month, year)
	m.MaritalStatus = jsonData.MaritalStatus[rand.IntN(len(jsonData.MaritalStatus))]
	m.Phone = fmt.Sprintf("156%d", rand.IntN(9999999))
	m.Email = fmt.Sprintf("%s%s%d@gmail.com", m.Name, m.LastName, year)
	m.PostalCode = strconv.Itoa(rand.IntN(8000) + 1000)
	m.Address = fmt.Sprintf("%s %d", jsonData.Streets[rand.IntN(len(jsonData.Streets))].Name, rand.IntN(9999))
	m.District = jsonData.Streets[rand.IntN(len(jsonData.Streets))].Name
	m.MemberNumber = strconv.Itoa(rand.IntN(9999999999))
	m.CUIL = fmt.Sprintf("%d-%s-%d", rand.IntN(9)+20, m.DNI, rand.IntN(8)+1)
	m.IdEnterprise = rand.IntN(500)
	fmt.Println(m)
	m.Category = jsonData.Categories[rand.IntN(len(jsonData.Categories))]
	entryYear := rand.IntN(64) + 1950
	entryMonth := rand.IntN(11) + 1
	var entryDay int
	switch month {
	case 2:
		entryDay = rand.IntN(27) + 1
	case 4, 6, 9, 11:
		entryDay = rand.IntN(29) + 1
	case 1, 3, 5, 7, 8, 10, 12:
		entryDay = rand.IntN(30) + 1
	}
	m.EntryDate = fmt.Sprintf("%d/%d/%d", entryDay, entryMonth, entryYear)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"member": m})
}
