package creators

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/models"
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
	type StreetData struct {
		Code string
		Name string
	}
	type JSONData struct {
		FemaleFirstNames []string
		MaleFirstNames   []string
		LastNames        []string
		MaritalStatus    []string
		Genders          []string
		Streets          []StreetData
	}

	m := models.Member{}
	file, err := os.Open("./data/jsonData.json")
	if err != nil {
		fmt.Println("error opening file")
		panic(err)
	}
	decoder := json.NewDecoder(file)
	jsonData := JSONData{}
	decoder.Decode(&jsonData)

	for i := 0; i < 100; i++ {
		// para que sea mas probable que sea hombre o mujer
		r := rand.IntN(6)
		if slices.Contains([]int{0, 1, 2}, r) {
			m.Gender = "Masculino"
		} else if slices.Contains([]int{3, 4, 5}, r) {
			m.Gender = "Hombre"
		} else {
			m.Gender = "Otro"
		}
		if m.Gender == "Masculino" {
			m.Name = jsonData.MaleFirstNames[rand.IntN(len(jsonData.MaleFirstNames))]
		} else if m.Gender == "Femenino" {
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
		// año entre 1950 y 2006 (mayor de 18 años)
		year := rand.IntN(56) + 1950
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
		// creo el dni con la fecha de nacimiento
		m.DNI = createDNI(day, month, year)
		dayStr := strconv.Itoa(day)
		monthStr := strconv.Itoa(month)
		yearStr := strconv.Itoa(year)
		if day < 10 {
			dayStr = "0" + dayStr
		}
		if month < 10 {
			monthStr = "0" + monthStr
		}

		// en la base de datos: '1998-05-22' string
		// consulta: SELECT > CAST('2023-06-26' AS DATE)
		// si lo quiero mostrar en el input lo doy vuelta y listo
		//
		// fijarse bien desp lo del formato fecha
		m.Birthday = fmt.Sprintf("%s/%s/%s", yearStr, monthStr, dayStr)
		// m.Birthday = time.Date(year, time.Month(month), day)
		m.MaritalStatus = jsonData.MaritalStatus[rand.IntN(len(jsonData.MaritalStatus))]
		m.Phone = fmt.Sprintf("156%d", rand.IntN(9999999))
		m.Email = fmt.Sprintf("%s%s%s@gmail.com", m.Name, m.LastName, strconv.Itoa(year)[2:])
		m.PostalCode = strconv.Itoa(rand.IntN(8000) + 1000)
		m.Address = fmt.Sprintf("%s %d", jsonData.Streets[rand.IntN(len(jsonData.Streets))].Name, rand.IntN(9999))
		m.District = jsonData.Streets[rand.IntN(len(jsonData.Streets))].Name
		m.MemberNumber = strconv.Itoa(rand.IntN(9999999999))
		m.CUIL = fmt.Sprintf("%d-%s-%d", rand.IntN(9)+20, m.DNI, rand.IntN(8)+1)
		m.IdEnterprise = rand.IntN(49) + 1

		switch {
		case year <= 1960:
			m.Category = "Nivel 1: Oficial Múltiple"
		case year <= 1970:
			m.Category = "Nivel 2: Oficial Especializado"
		case year <= 1980:
			m.Category = "Nivel 3: Oficial General"
		case year <= 1990:
			m.Category = "Nivel 4: Medio Oficial"
		case year <= 2000:
			m.Category = "Nivel 5: Ayudante"
		case year <= 2006:
			m.Category = "Nivel 6: Operario Act. Industrial"
		}

		// que sea a los 18 años o mas, entre 18 y 48
		entryYear := rand.IntN(30) + year + 18
		if entryYear > 2024 {
			entryYear = 2024
		}
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
		entryDayStr := strconv.Itoa(entryDay)
		entryMonthStr := strconv.Itoa(entryMonth)
		entryYearStr := strconv.Itoa(entryYear)
		if entryDay < 10 {
			entryDayStr = "0" + entryDayStr
		}
		if entryMonth < 10 {
			entryMonthStr = "0" + entryMonthStr
		}
		m.EntryDate = fmt.Sprintf("%s/%s/%s", entryYearStr, entryMonthStr, entryDayStr)
		insert, err := database.DB.Query("INSERT INTO MemberTable (Name, LastName, DNI, Birthday, Gender, MaritalStatus, Phone, Email, Address, PostalCode, District, MemberNumber, CUIL, IdEnterprise, Category, EntryDate) VALUES ('?','?','?','?','?','?','?','?','?','?','?','?','?','?','?','?')", m.Name, m.LastName, m.DNI, m.Birthday, m.Gender, m.MaritalStatus, m.Phone, m.Email, m.Address, m.PostalCode, m.District, m.MemberNumber, m.CUIL, m.IdEnterprise, m.Category, m.EntryDate)
		if err != nil {
			fmt.Println("error inserting member")
			panic(err)
		}
		insert.Close()
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "members hecho"})
}
