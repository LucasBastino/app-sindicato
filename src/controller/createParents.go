package controller

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

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
		// fijarse lo de fecha
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
	p.IdMember = rand.IntN(1000)
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
	// hacer el query
	p.LastName = jsonData.LastNames[rand.IntN(len(jsonData.LastNames))]
	// m.Birthday = fmt.Sprintf("%d/%d/%d", day, month, year)
	// m.CUIL = fmt.Sprintf("%d-%s-%d", rand.IntN(9)+20, m.DNI, rand.IntN(8)+1)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"parent": p})
}
