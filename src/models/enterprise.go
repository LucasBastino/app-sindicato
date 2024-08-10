package models

import (
	"fmt"
	"strings"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Enterprise struct {
	IdEnterprise int
	Name         string
	Address      string
}

func (e Enterprise) InsertModel() {
	insert, err := database.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s','%s')", e.Name, e.Address))
	if err != nil {
		// DBError{"INSERT Enterprise"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	defer insert.Close()
}

func (e Enterprise) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf("DELETE FROM EnterpriseTable WHERE IdEnterprise = '%v'", e.IdEnterprise))
	if err != nil {
		// DBError{"DELETE Enterprise"}.Error(err)
		fmt.Println("error deleting Enterprise")
		panic(err)
	}
	defer delete.Close()

}

func (e Enterprise) EditModel() {
	update, err := database.DB.Query(fmt.Sprintf("UPDATE EnterpriseTable SET Name = '%s', Address = '%s' WHERE IdEnterprise = '%v'", e.Name, e.Address, e.IdEnterprise))
	if err != nil {
		// DBError{"UPDATE Enterprise"}.Error(err)
		fmt.Println("error updating Enterprise")
		panic(err)
	}
	defer update.Close()
}

func (e Enterprise) GetIdModel(c *fiber.Ctx) int {
	params := struct {
		IdEnterprise int `params:"IdEnterprise"`
	}{}

	c.ParamsParser(&params)

	return params.IdEnterprise

	// IdEnterpriseStr := c.PathValue("IdEnterprise")
	// IdEnterprise, err := strconv.Atoi(IdEnterpriseStr)
	// if err != nil {
	// 	fmt.Println("error converting type")
	// 	panic(err)
	// }
	// return IdEnterprise
}

func (e Enterprise) SearchOneModelById(c *fiber.Ctx) Enterprise {
	IdEnterprise := e.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf("SELECT IdEnterprise, Name, Address FROM EnterpriseTable WHERE IdEnterprise = '%v'", IdEnterprise))
	if err != nil {
		fmt.Println("error searching enterprise by id")
		panic(err)
	}

	var enterprise Enterprise
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning Enterprise")
			panic(err)
		}
	}
	defer result.Close()
	return enterprise
}

func (e Enterprise) SearchModels(c *fiber.Ctx) []Enterprise {
	searchKey := c.FormValue("search-key")
	var enterprises []Enterprise
	var enterprise Enterprise

	result, err := database.DB.Query(fmt.Sprintf(`SELECT * FROM EnterpriseTable WHERE Name LIKE '%%%s%%' OR Address LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching Enterprise in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&enterprise.IdEnterprise, &enterprise.Name, &enterprise.Address)
		if err != nil {
			fmt.Println("error scanning Enterprise")
			panic(err)
		}
		enterprises = append(enterprises, enterprise)
	}
	defer result.Close()
	return enterprises
}

func (e Enterprise) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}
	if strings.TrimSpace(c.FormValue("name")) == "" {
		errorMap["name"] = "el campo Nombre no puede estar vacio"
	}
	if strings.TrimSpace(c.FormValue("address")) == "" {
		errorMap["address"] = "el campo Direccion no puede estar vacio"
	}
	return errorMap
}

func (e Enterprise) CreateTemplateData(enterprise Enterprise, enterprises []Enterprise, path string, errorMap map[string]string) TemplateData {
	templateData := TemplateData{}
	templateData.Enterprise = enterprise
	templateData.Enterprises = enterprises
	templateData.Path = path
	templateData.ErrorMap = errorMap
	return templateData
}
