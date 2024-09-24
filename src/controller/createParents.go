package controller

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

// func CreateParents(c *fiber.Ctx) error {
// 	var Name string
// 	var Rel string
// 	var IdMember int
// 	var IdMemberArray []int
// 	result, err := database.DB.Query("SELECT IdMember from MemberTable")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	for result.Next() {
// 		err = result.Scan(&IdMember)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		IdMemberArray = append(IdMemberArray, IdMember)
// 	}

// 	for _, IdMember := range IdMemberArray {
// 		for i := 1; i <= 4; i++ {
// 			Name = fmt.Sprintf("pariente%d", rand.IntN(99)+1)
// 			Rel = fmt.Sprintf("rel%d", IdMember)
// 			insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('%s','%s', '%d')", Name, Rel, IdMember))
// 			if err != nil {
// 				fmt.Println("error insertando en la DB")
// 				panic(err)
// 			}
// 			insert.Close()

// 		}

// 	}
// 	return c.SendString("parents created")
// }

func CreateParents(c *fiber.Ctx) error {
	type Parent struct {
		IdParent int
		Name     string
		LastName string
		Rel      string
		Birthday string
		Gender   string
		CUIL     string
		IdMember int
	}
	type JSONData struct {
		MaleFirstNames   []string
		FemaleFirstNames []string
		LastNames        []string
		Genders          []string
	}

	p := Parent{}
	file, err := os.Open("./data/jsonData.json")
	if err != nil {
		fmt.Println("error opening file")
		panic(err)
	}
	decoder := json.NewDecoder(file)
	jsonData := JSONData{}
	decoder.Decode(&jsonData)
	p.IdParent = rand.IntN(3000)
	p.IdMember = rand.IntN(200)
	p.Gender = jsonData.Genders[rand.IntN(len(jsonData.Genders))]
	if p.Gender == "Hombre" {
		p.Name = jsonData.MaleFirstNames[rand.IntN(len(jsonData.MaleFirstNames))]
		p.Rel = "Hijo"
	} else if p.Gender == "Mujer" {
		p.Name = jsonData.FemaleFirstNames[rand.IntN(len(jsonData.FemaleFirstNames))]
		p.Rel = "Hija"
	} else if p.Gender == "Otro" {
		p.Rel = "Hijx"
		random := rand.IntN(1)
		if random == 0 {
			p.Name = jsonData.MaleFirstNames[rand.IntN(len(jsonData.MaleFirstNames))]
		} else {
			p.Name = jsonData.FemaleFirstNames[rand.IntN(len(jsonData.FemaleFirstNames))]
		}
	}
	p.LastName = jsonData.LastNames[rand.IntN(len(jsonData.LastNames))]
	// BUSCAR LA MANERA DE QUE SOLO TE APAREZCAN LOS MENORES DE 50
	// con birthday formato fecha y una funcion en SQL que saque el año y que busque que sea mayor a 1974
	// tambien hacer que el form de fechas sea con 3 input text y que despues de ahi se pase a fecha en backend
	// y que para mostrarlo se obtenga de la DB en formato fecha, en backend se pase a string y de ahi se renderiza en los 3 inputs
	// o en 1 input pero separados con barras (no se como se hace)
	result, err := database.DB.Query(fmt.Sprintf("SELECT Birthday FROM MemberTable WHERE IdMember = '%d'", p.IdMember))
	if err != nil {
		fmt.Println("error consulting member table")
		panic(err)
	}
	var memberBirthday string
	for result.Next() {
		err = result.Scan(&memberBirthday)
		if err != nil {
			fmt.Println("error scanning member birthday")
			panic(err)
		}
	}
	result.Close()
	// consulto la fecha de nacimiento del afiliado mayor y obtengo el año
	memberYearStr := memberBirthday[(len(memberBirthday) - 4):]
	memberYear, err := strconv.Atoi(memberYearStr)
	if err != nil {
		fmt.Println("error converting memberYearStr to int")
		panic(err)
	}
	fmt.Println(memberYear)
	// la edad del hijo sera entre 0 y 25 años, por lo tanto va a nacer entre 1999 y 2024
	year := rand.IntN(25) + 1999
	ageDifference := year - memberYear
	if ageDifference < 20 {
		year = year + 20 - ageDifference
	}
	if year > 2024 {
		year = 2024
	} else if year < 1999 {
		year = 1999
	}
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
	p.Birthday = fmt.Sprintf("%d/%d/%d", day, month, year)
	p.CUIL = fmt.Sprintf("%d-%s-%d", rand.IntN(9)+20, strconv.Itoa(rand.IntN(3000000)+2000000), rand.IntN(8)+1)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"parent": p})
}
