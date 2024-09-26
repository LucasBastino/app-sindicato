package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/models"
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
	type JSONData struct {
		MaleFirstNames   []string
		FemaleFirstNames []string
		LastNames        []string
		Genders          []string
	}

	p := models.Parent{}

	file, err := os.Open("./data/jsonData.json")
	if err != nil {
		fmt.Println("error opening file")
		panic(err)
	}
	decoder := json.NewDecoder(file)
	jsonData := JSONData{}
	decoder.Decode(&jsonData)

	// BUSCAR LA MANERA DE QUE SOLO TE APAREZCAN LOS MENORES DE 50
	// con birthday formato fecha y una funcion en SQL que saque el año y que busque que sea mayor a 1974
	// tambien hacer que el form de fechas sea con 3 input text y que despues de ahi se pase a fecha en backend
	// y que para mostrarlo se obtenga de la DB en formato fecha, en backend se pase a string y de ahi se renderiza en los 3 inputs
	// o en 1 input pero separados con barras (no se como se hace)

	// busco los afiliados menores de 50 años, esos van a tener hijos
	result, err := database.DB.Query("SELECT IdMember, LastName, Birthday FROM MemberTable WHERE CAST(Birthday AS DATE) > '1974-01-01'")
	if err != nil {
		fmt.Println("error consulting member table")
		panic(err)
	}

	// creo un struct para facilitar la lectura de datos
	type MemberDateInfo struct {
		id       int
		lastName string
		birthday string
	}
	var memberDateInfoList []MemberDateInfo
	var memberDateInfo MemberDateInfo

	// guardo los datos de cada afiliado en un slice
	for result.Next() {
		err = result.Scan(&memberDateInfo.id, &memberDateInfo.lastName, &memberDateInfo.birthday)
		if err != nil {
			fmt.Println("error scanning member id and birthday")
			panic(err)
		}
		memberDateInfoList = append(memberDateInfoList, memberDateInfo)
	}
	result.Close()

	// for i := 0; i < 500; i++ {
	randomMember := memberDateInfoList[rand.IntN(len(memberDateInfoList))]
	// obtengo el año de nacimiento del afiliado mayor
	memberYearStr := randomMember.birthday[:4]
	memberYear, err := strconv.Atoi(memberYearStr)
	if err != nil {
		fmt.Println("error converting memberYearStr to int")
		panic(err)
	}
	birthday, dni := createBirthdayAndDNI(memberYear)
	p.Birthday = birthday
	p.CUIL = dni
	p.IdMember = randomMember.id

	p.CUIL = fmt.Sprintf("%d-%s-%d", rand.IntN(9)+20, strconv.Itoa(rand.IntN(3000000)+2000000), rand.IntN(8)+1)
	// para que sea mas probable que sea hombre o mujer
	r := rand.IntN(6)
	if slices.Contains([]int{0, 1, 2}, r) {
		p.Gender = "Hombre"
	} else if slices.Contains([]int{3, 4, 5}, r) {
		p.Gender = "Mujer"
	} else {
		p.Gender = "Otro"
	}

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
	p.LastName = randomMember.lastName

	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, LastName, Rel, Birthday, Gender, CUIL, IdMember) VALUES ('%s','%s', '%s', '%s', '%s', '%s', '%d')", p.Name, p.LastName, p.Rel, p.Birthday, p.Gender, p.CUIL, p.IdMember))
	if err != nil {
		fmt.Println("error inserting parent")
		panic(err)
	}
	insert.Close()
	// }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": p})

}

func createBirthdayAndDNI(memberYear int) (string, string) {
	// la edad del hijo sera entre 0 y 25 años, por lo tanto va a nacer entre 1999 y 2024
	year := rand.IntN(25) + 1998
	ageDifference := year - memberYear
	if ageDifference < 20 {
		year = year + 20 - ageDifference
	}
	if year > 2023 {
		year = 2023
	} else if year < 1998 {
		year = 1998
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
	birthdayStr := fmt.Sprintf("%d-%d-%d", year, month, day)
	birthdayDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	diff := math.Round(birthdayDate.Sub(time.Date(1900, 01, 01, 0, 0, 0, 0, time.UTC)).Hours())

	// creador de DNI
	// funcion hecha con chatGPT con los siguientes valores
	// 650376 horas con dni 23885185, 13 marzo 1974
	// 845016 horas con dni 39713471, 26 mayo 1996
	// 903576 horas con dni 44594659 enero 2003

	dniInt := int((81.21 * diff) - 28986071.36)
	dni := strconv.Itoa(dniInt)
	return birthdayStr, dni
}
