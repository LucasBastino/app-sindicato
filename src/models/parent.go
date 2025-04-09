package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
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

func (parent Parent) InsertModel() (Parent, error) {
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
		er.QueryError.Msg = err.Error()
		return Parent{}, er.QueryError
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT * FROM ParentTable 
		WHERE IdParent = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Parent{}, er.QueryError
	}
	p, _, err := parent.ScanResult(result, true)
	if err != nil {
		return Parent{}, err
	}
	return p, nil
}

func (parent Parent) DeleteModel() error {
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM ParentTable 
		WHERE IdParent = '%d'`,
		parent.IdParent))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return er.QueryError
	}
	defer delete.Close()
	return nil
}

func (parent Parent) UpdateModel() (Parent, error) {
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
		er.QueryError.Msg = err.Error()
		return Parent{}, er.QueryError
	}
	update.Close()
	result, err := database.DB.Query(fmt.Sprintf(`
	SELECT * FROM ParentTable WHERE IdParent = '%d'`, parent.IdParent))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Parent{}, er.QueryError
	}
	p, _, err := parent.ScanResult(result, true)
	if err != nil {
		return Parent{}, err
	}
	return p, nil
}

func (parent Parent) GetIdModel(c *fiber.Ctx) (int, error) {
	params := struct {
		IdParent int `params:"IdParent"`
	}{}
	err := c.ParamsParser(&params)
	if err != nil {
		er.ParamsError.Msg = err.Error()
		return 0, er.ParamsError
	}
	return params.IdParent, nil
}

func (parent Parent) SearchOneModelById(c *fiber.Ctx) (Parent, error) {
	IdParent, err := parent.GetIdModel(c)
	if err != nil {
		return Parent{}, err
	}
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM ParentTable
		WHERE IdParent = '%d'`, IdParent))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Parent{}, er.QueryError
	}
	p, _, err := parent.ScanResult(result, true)
	if err != nil {
		return Parent{}, err
	}
	return p, nil
}

func (parent Parent) SearchModels(c *fiber.Ctx, offset int) ([]Parent, string, error) {
	idMember, err := Member{}.GetIdModel(c)
	if err != nil {
		return nil, "", err
	}
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM ParentTable 
		WHERE IdMember = %d`, idMember))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, "", er.QueryError
	}
	_, pp, err := parent.ScanResult(result, false)
	if err != nil {
		return nil, "", err
	}
	return pp, "", nil
}

func (parent Parent) ValidateFields(c *fiber.Ctx) (map[string]string, error) {
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
	if len(errorMap) > 1 {

		return errorMap, er.ValidationError
	}
	return errorMap, nil
}

func (parent Parent) GetTotalRows(c *fiber.Ctx) (int, error) {
	var totalRows int
	idMember, err := Member{}.GetIdModel(c)
	if err != nil {
		return 0, err
	}
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE IdMember = '%d'`, idMember))
	// row.Scan copia el numero de fila en la variable count
	err = row.Scan(&totalRows)
	if err != nil {
		er.ScanError.Msg = err.Error()
		return 0, er.ScanError
	}
	return totalRows, nil
}

func (parent Parent) GetFiberMap(parents []Parent, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return nil
}

func (parent Parent) GetAllModels() []Parent {
	return nil
}

func (parent Parent) ScanResult(result *sql.Rows, onlyOne bool) (Parent, []Parent, error) {
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
			er.ScanError.Msg = err.Error()
			return Parent{}, nil, er.ScanError
		}
		if !onlyOne {
			pp = append(pp, p)
		}
	}
	result.Close()
	return p, pp, nil
}

func (parent Parent) CheckDeleted(idParent int) (bool, error) {
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
		er.ScanError.Msg = err.Error()
		return false, er.ScanError
	}
	if totalRows == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
