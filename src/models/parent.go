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

func (parent Parent) InsertModel() Parent {
	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('%s','%s', '%d')", parent.Name, parent.Rel, parent.IdMember))
	if err != nil {
		// DBError{"INSERT Parent"}.Error(err)
		fmt.Println("error inserting parent")
		panic(err)
	}
	insert.Close()
	var p Parent
	result, err := database.DB.Query("SELECT * FROM ParentTable WHERE IdParent = (SELECT LAST_INSERT_ID())")
	if err != nil {
		fmt.Print(err)
	}
	result.Next()
	err = result.Scan(&p.IdParent, &p.Name, &p.Rel, &p.IdMember)
	if err != nil {
		fmt.Println(err)
	}
	result.Close()
	return p
}

func (parent Parent) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf("DELETE FROM ParentTable WHERE IdParent = '%v'", parent.IdParent))
	if err != nil {
		// DBError{"DELETE Parent"}.Error(err)
		fmt.Println("error deleting parent")
		panic(err)
	}
	defer delete.Close()

}

func (parent Parent) EditModel() {
	update, err := database.DB.Query(fmt.Sprintf("UPDATE ParentTable SET Name = '%s', Rel = '%s' WHERE IdParent = '%v'", parent.Name, parent.Rel, parent.IdParent))
	if err != nil {
		// DBError{"UPDATE Parent"}.Error(err)
		fmt.Println("error updating parent")
		panic(err)
	}
	defer update.Close()
}

func (parent Parent) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdParent int `params:"IdParent"`
	}{}
	c.ParamsParser(&params)
	return params.IdParent
}

func (parent Parent) SearchOneModelById(c *fiber.Ctx) Parent {
	IdParent := parent.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf("SELECT IdParent, Name, Rel, IdMember FROM ParentTable WHERE IdParent = '%v'", IdParent))
	if err != nil {
		fmt.Println("error searching parent by id")
		panic(err)
	}

	var p Parent
	for result.Next() {
		err = result.Scan(&p.IdParent, &p.Name, &p.Rel, &p.IdMember)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
	}
	defer result.Close()
	return p
}

func (parent Parent) SearchModels(c *fiber.Ctx, offset int) ([]Parent, string) {
	searchKey := c.FormValue("search-key")
	var parents []Parent
	var p Parent

	result, err := database.DB.Query(fmt.Sprintf(`SELECT IdParent, Name, Rel FROM ParentTable WHERE Name LIKE '%%%s%%' OR Rel LIKE '%%%s%%' LIMIT 10 OFFSET %d`, searchKey, searchKey, offset))
	if err != nil {
		fmt.Println("error searching Parent in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&p.IdParent, &p.Name, &p.Rel)
		if err != nil {
			fmt.Println("error scanning Parent")
			panic(err)
		}
		parents = append(parents, p)
	}
	defer result.Close()
	return parents, searchKey
}

func (parent Parent) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}
	if strings.TrimSpace(c.FormValue("name")) == "" {
		errorMap["name"] = "el campo Nombre no puede estar vacio"
	}
	if strings.TrimSpace(c.FormValue("rel")) == "" {
		errorMap["rel"] = "el campo Parentesco no puede estar vacio"
	}
	return errorMap
}

func (parent Parent) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	searchKey := c.FormValue("search-key")
	row := database.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM ParentTable WHERE Name LIKE '%%%s%%'", searchKey))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	return totalRows
}

func (parent Parent) GetFiberMap(parents []Parent, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return fiber.Map{
		"model":           "parent",
		"parents":         parents,
		"searchKey":       searchKey,
		"currentPage":     currentPage,
		"firstPage":       1,
		"previousPage":    currentPage - 1,
		"someBefore":      currentPage - someBefore,
		"sixBefore":       currentPage - 6,
		"fiveBefore":      currentPage - 5,
		"fourBefore":      currentPage - 4,
		"threeBefore":     currentPage - 3,
		"twoBefore":       currentPage - 2,
		"twoAfter":        currentPage + 2,
		"threeAfter":      currentPage + 3,
		"fourAfter":       currentPage + 4,
		"fiveAfter":       currentPage + 5,
		"sixAfter":        currentPage + 6,
		"someAfter":       currentPage + someAfter,
		"nextPage":        currentPage + 1,
		"lastPage":        totalPages,
		"totalPages":      totalPages,
		"totalPagesArray": totalPagesArray,
	}
}

func (parent Parent) GetAllModels() []Parent {
	var parents []Parent
	var p Parent

	result, err := database.DB.Query("SELECT * FROM ParentTable")
	if err != nil {
		fmt.Println("error searching parent in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&p.IdParent, &p.Name, &p.Rel, &p.IdMember)
		if err != nil {
			fmt.Println("error scanning enterprise")
			panic(err)
		}
		parents = append(parents, p)
	}
	defer result.Close()
	return parents
}
