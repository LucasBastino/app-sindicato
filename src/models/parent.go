package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Parent struct {
	IdParent int
	Name     string
	Rel      string
	IdMember int
}

func (newParent Parent) InsertModel() {
	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('%s','%s', '%d')", newParent.Name, newParent.Rel, newParent.IdMember))
	if err != nil {
		// DBError{"INSERT Parent"}.Error(err)
		fmt.Println("error inserting parent")
		panic(err)
	}
	defer insert.Close()
}

func (p Parent) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf("DELETE FROM ParentTable WHERE IdParent = '%v'", p.IdParent))
	if err != nil {
		// DBError{"DELETE Parent"}.Error(err)
		fmt.Println("error deleting parent")
		panic(err)
	}
	defer delete.Close()

}

func (p Parent) EditModel() {
	update, err := database.DB.Query(fmt.Sprintf("UPDATE ParentTable SET Name = '%s', Rel = '%s' WHERE IdParent = '%v'", p.Name, p.Rel, p.IdParent))
	if err != nil {
		// DBError{"UPDATE Parent"}.Error(err)
		fmt.Println("error updating parent")
		panic(err)
	}
	defer update.Close()
}

func (p Parent) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdParent int `params:"IdParent"`
	}{}
	c.ParamsParser(&params)
	return params.IdParent
}

func (p Parent) SearchOneModelById(c *fiber.Ctx) Parent {
	IdParent := p.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf("SELECT IdParent, Name, Rel, IdMember FROM ParentTable WHERE IdParent = '%v'", IdParent))
	if err != nil {
		fmt.Println("error searching parent by id")
		panic(err)
	}

	var parent Parent
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel, &parent.IdMember)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
	}
	defer result.Close()
	return parent
}

func (p Parent) SearchModels(c *fiber.Ctx, offset int) []Parent {
	searchKey := c.FormValue("search-key")
	var parents []Parent
	var parent Parent

	result, err := database.DB.Query(fmt.Sprintf(`SELECT IdParent, Name, Rel FROM ParentTable WHERE Name LIKE '%%%s%%' OR Rel LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching Parent in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&parent.IdParent, &parent.Name, &parent.Rel)
		if err != nil {
			fmt.Println("error scanning Parent")
			panic(err)
		}
		parents = append(parents, parent)
	}
	defer result.Close()
	return parents
}

func (p Parent) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}
	if strings.TrimSpace(c.FormValue("name")) == "" {
		errorMap["name"] = "el campo Nombre no puede estar vacio"
	}
	if strings.TrimSpace(c.FormValue("rel")) == "" {
		errorMap["rel"] = "el campo Parentesco no puede estar vacio"
	}
	return errorMap
}

func (p Parent) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow("SELECT COUNT(*) FROM ParentTable WHERE Name LIKE '%%%s%%'", searchKey)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	return totalRows
}
