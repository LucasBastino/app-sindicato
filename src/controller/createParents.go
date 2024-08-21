package controller

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

func CreateParents(c *fiber.Ctx) error {
	var Name string
	var Rel string
	var IdMember int
	var IdMemberArray []int
	result, err := database.DB.Query("SELECT IdMember from MemberTable")
	if err != nil {
		log.Panic(err)
	}
	for result.Next() {
		err = result.Scan(&IdMember)
		if err != nil {
			log.Panic(err)
		}
		IdMemberArray = append(IdMemberArray, IdMember)
	}

	for _, IdMember := range IdMemberArray {
		for i := 1; i <= 5; i++ {
			Name = fmt.Sprintf("pariente%d", rand.IntN(99)+1)
			Rel = fmt.Sprintf("rel%d", IdMember)
			insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('%s','%s', '%d')", Name, Rel, IdMember))
			if err != nil {
				fmt.Println("error insertando en la DB")
				panic(err)
			}
			insert.Close()

		}

	}
	return c.SendString("parents created")
}
