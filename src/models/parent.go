package models

import (
	"database/sql"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/gofiber/fiber/v2"
)

type Parent struct {
	IdParent  int		`json:"idparent"`
	Name      string	`json:"name"`
	LastName  string	`json:"lastname"`
	Rel       string	`json:"rel"`
	Birthday  string	`json:"birthday"`
	Gender    string	`json:"gender"`
	CUIL      string	`json:"cuil"`
	IdMember  int		`json:"idmember"`
	CreatedAt time.Time	`json:"createdat"`
	UpdatedAt time.Time	`json:"updatedat"`
}

func (parent Parent) InsertModel() (Parent, customError.CustomError) {
	parent.Birthday = FormatToYYYYMMDD(parent.Birthday)
	insert, err := database.DB.Query(`
		INSERT INTO ParentTable 
		(Name,
		LastName,
		Rel,
		Birthday,
		Gender,
		CUIL,
		IdMember)
		VALUES (?,?,?, ?, ?, ?, ?)`,
		parent.Name,
		parent.LastName,
		parent.Rel,
		parent.Birthday,
		parent.Gender,
		parent.CUIL,
		parent.IdMember)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Parent{}, customError.QueryError
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT * FROM ParentTable 
		WHERE IdParent = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Parent{}, customError.QueryError
	}
	p, _, customErr := parent.ScanResult(result, true)
	if (customErr != customError.CustomError{}) {
		return Parent{}, customErr
	}
	return p, customError.CustomError{}
}

func (parent Parent) DeleteModel() customError.CustomError {
	delete, err := database.DB.Query(`
		DELETE FROM ParentTable 
		WHERE IdParent = ?`,
		parent.IdParent)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return customError.QueryError
	}
	defer delete.Close()
	return customError.CustomError{}
}

func (parent Parent) UpdateModel() (Parent, customError.CustomError) {
	parent.Birthday = FormatToYYYYMMDD(parent.Birthday)
	update, err := database.DB.Query(`
		UPDATE ParentTable 
		SET Name = ?,
		LastName = ?,
		Rel = ?,
		Birthday = ?,
		Gender = ?,
		CUIL = ?,
		IdMember = ?
		WHERE IdParent = ?`,
		parent.Name,
		parent.LastName,
		parent.Rel,
		parent.Birthday,
		parent.Gender,
		parent.CUIL,
		parent.IdMember,
		parent.IdParent)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Parent{}, customError.QueryError
	}
	update.Close()
	result, err := database.DB.Query(`
	SELECT * FROM ParentTable WHERE IdParent = ?`, parent.IdParent)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Parent{}, customError.QueryError
	}
	p, _, customErr := parent.ScanResult(result, true)
	if (err != customError.CustomError{}) {
		return Parent{}, customErr
	}
	return p, customError.CustomError{}
}

func (parent Parent) GetIdModel(c *fiber.Ctx) (int, customError.CustomError) {
	params := struct {
		IdParent int `params:"IdParent"`
	}{}
	err := c.ParamsParser(&params)
	if err != nil {
		customError.ParamsError.Msg = err.Error()
		return 0, customError.ParamsError
	}
	return params.IdParent, customError.CustomError{}
}

func (parent Parent) SearchOneModelById(c *fiber.Ctx) (Parent, customError.CustomError) {
	IdParent, customErr := parent.GetIdModel(c)
	if (customErr != customError.CustomError{}) {
		return Parent{}, customErr
	}
	result, err := database.DB.Query(`
		SELECT
		*
		FROM ParentTable
		WHERE IdParent = ?`, IdParent)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return Parent{}, customError.QueryError
	}
	p, _, customErr := parent.ScanResult(result, true)
	if (customErr != customError.CustomError{}) {
		return Parent{}, customErr
	}
	return p, customError.CustomError{}
}

func (parent Parent) SearchModels(c *fiber.Ctx, offset int) ([]Parent, string, customError.CustomError) {
	idMember, customErr := Member{}.GetIdModel(c)
	if (customErr != customError.CustomError{}) {
		return nil, "", customErr
	}
	result, err := database.DB.Query(`
		SELECT
		*
		FROM ParentTable 
		WHERE IdMember = ?`, idMember)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return nil, "", customError.QueryError
	}
	_, pp, customErr := parent.ScanResult(result, false)
	if (customErr != customError.CustomError{}) {
		return nil, "", customErr
	}
	return pp, "", customError.CustomError{}
}

func (parent Parent) ValidateFields(c *fiber.Ctx) customError.CustomError {
	validateFunctions := []func(*fiber.Ctx) error{
		ValidateName,
		ValidateLastName,
		ValidateRel,
		ValidateBirthday,
		ValidateGender,
		ValidateCUIL,
	}
	for _, vF := range validateFunctions {
		if err := vF(c); err != nil {
			customError.ValidationError.Msg = err.Error()
			return customError.ValidationError
		} else {
			continue
		}
	}
	return customError.CustomError{}
}

func (parent Parent) GetTotalRows(c *fiber.Ctx) (int, customError.CustomError) {
	var totalRows int
	idMember, customErr := Member{}.GetIdModel(c)
	if (customErr != customError.CustomError{}) {
		return 0, customErr
	}
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE IdMember = ?`, idMember)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		customError.ScanError.Msg = err.Error()
		return 0, customError.ScanError
	}
	return totalRows, customError.CustomError{}
}

func (parent Parent) GetFiberMap(parents []Parent, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return nil
}

func (parent Parent) GetAllModels() ([]Parent, customError.CustomError) {
	return nil, customError.CustomError{}
}

func (parent Parent) ScanResult(result *sql.Rows, onlyOne bool) (Parent, []Parent, customError.CustomError) {
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
			customError.ScanError.Msg = err.Error()
			return Parent{}, nil, customError.ScanError
		}
		if !onlyOne {
			pp = append(pp, p)
		}
	}
	result.Close()
	return p, pp, customError.CustomError{}
}

func (parent Parent) CheckDeleted(idParent int) (bool, customError.CustomError) {
	var totalRows int
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM ParentTable
	// 	WHERE IdParent = '%d'`, parent.IdParent))
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE IdParent = ?`, idParent)
	// row.Scan copia el numero de fila en la variable count
	err := row.Scan(&totalRows)
	if err != nil {
		customError.ScanError.Msg = err.Error()
		return false, customError.ScanError
	}
	if totalRows == 0 {
		return true, customError.CustomError{}
	} else {
		return false, customError.CustomError{}
	}
}
