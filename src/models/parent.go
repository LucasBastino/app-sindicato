package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/gofiber/fiber/v2"
)

type Parent struct {
	IdParent  int
	Name      string
	LastName  string
	Rel       string
	Birthday  string
	Gender    string
	CUIL      string
	IdMember  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (parent Parent) InsertModel() Parent {
	parent.Birthday = FormatToYYYYMMDD(parent.Birthday)
	insert, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO ParentTable 
		(Name,
		LastName,
		Rel,
		Birthday,
		Gender,
		CUIL,
		IdMember)
		VALUES ('%s','%s','%s', '%s', '%s', '%s', '%d')`,
		parent.Name,
		parent.LastName,
		parent.Rel,
		parent.Birthday,
		parent.Gender,
		parent.CUIL,
		parent.IdMember))
	if err != nil {
		// DBError{"INSERT Parent"}.Error(err)
		fmt.Println("error inserting parent")
		panic(err)
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT * FROM ParentTable 
		WHERE IdParent = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		fmt.Print(err)
	}
	p, _ := parent.ScanResult(result, true)
	return p
}

func (parent Parent) DeleteModel() {
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM ParentTable 
		WHERE IdParent = '%d'`,
		parent.IdParent))
	if err != nil {
		// DBError{"DELETE Parent"}.Error(err)
		fmt.Println("error deleting parent")
		panic(err)
	}
	defer delete.Close()

}

func (parent Parent) UpdateModel() Parent {
	parent.Birthday = FormatToYYYYMMDD(parent.Birthday)
	update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE ParentTable 
		SET Name = '%s',
		LastName = '%s',
		Rel = '%s',
		Birthday = '%s',
		Gender = '%s',
		CUIL = '%s',
		IdMember = '%d'
		WHERE IdParent = '%d'`,
		parent.Name,
		parent.LastName,
		parent.Rel,
		parent.Birthday,
		parent.Gender,
		parent.CUIL,
		parent.IdMember,
		parent.IdParent))
	if err != nil {
		// DBError{"UPDATE Parent"}.Error(err)
		fmt.Println("error updating parent")
		panic(err)
	}
	update.Close()
	result, err := database.DB.Query(`
		SELECT * FROM ParentTable 
		WHERE IdParent = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		fmt.Print(err)
	}
	p, _ := parent.ScanResult(result, true)
	return p
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
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM ParentTable
		WHERE IdParent = '%d'`, IdParent))
	if err != nil {
		fmt.Println("error searching parent by id")
		panic(err)
	}
	p, _ := parent.ScanResult(result, true)
	return p
}

func (parent Parent) SearchModels(c *fiber.Ctx, offset int) ([]Parent, string) {
	idMember := Member{}.GetIdModel(c)
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM ParentTable 
		WHERE IdMember = %d`, idMember))
	if err != nil {
		fmt.Println("error searching member parents in DB")
		panic(err)
	}
	_, pp := parent.ScanResult(result, false)
	return pp, ""
}

func (parent Parent) ValidateFields(c *fiber.Ctx) map[string]string {
	errorMap := map[string]string{}

	var valid bool
	var err string

	if valid, err = ValidateName(c); !valid {
		errorMap["name"] = err
	}
	if valid, err = ValidateLastName(c); !valid {
		errorMap["lastName"] = err
	}
	if valid, err = ValidateRel(c); !valid {
		errorMap["rel"] = err
	}
	if valid, err = ValidateBirthday(c); !valid {
		errorMap["birthday"] = err
	}
	if valid, err = ValidateGender(c); !valid {
		errorMap["gender"] = err
	}
	if valid, err = ValidateCUIL(c); !valid {
		errorMap["cuil"] = err
	}
	return errorMap
}

func (parent Parent) GetTotalRows(c *fiber.Ctx) int {
	var totalRows int
	idMember := Member{}.GetIdModel(c)
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE IdMember = '%d'`, idMember))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	return totalRows
}

func (parent Parent) GetFiberMap(parents []Parent, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return nil
}

func (parent Parent) GetAllModels() []Parent {
	return nil
}

func (parent Parent) ScanResult(result *sql.Rows, onlyOne bool) (Parent, []Parent) {
	var p Parent
	var pp []Parent
	for result.Next() {
		err := result.Scan(
			&p.IdParent,
			&p.Name,
			&p.LastName,
			&p.Rel,
			&p.Birthday,
			&p.Gender,
			&p.CUIL,
			&p.IdMember,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		// formateo las fechas en formato argentino
		p.Birthday = FormatToDDMMYYYY(p.Birthday)
		if err != nil {
			fmt.Println("error scanning parent")
			panic(err)
		}
		if !onlyOne {
			pp = append(pp, p)
		}
	}
	result.Close()
	return p, pp
}

func (parent Parent) CheckDeleted(idParent int) bool {
	var totalRows int
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM ParentTable
	// 	WHERE IdParent = '%d'`, parent.IdParent))
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE IdParent = '%d'`, idParent))
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		log.Fatal(err)
	}
	if totalRows == 0 {
		return true
	} else {
		return false
	}
}
