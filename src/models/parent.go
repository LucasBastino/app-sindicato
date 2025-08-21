package models

import (
	"database/sql"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
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

func (parent Parent) InsertModel() (Parent, error) {
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
	delete, err := database.DB.Query(`
		DELETE FROM ParentTable 
		WHERE IdParent = ?`,
		parent.IdParent)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return er.QueryError
	}
	defer delete.Close()
	return nil
}

func (parent Parent) UpdateModel() (Parent, error) {
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
		er.QueryError.Msg = err.Error()
		return Parent{}, er.QueryError
	}
	update.Close()
	result, err := database.DB.Query(`
	SELECT * FROM ParentTable WHERE IdParent = ?`, parent.IdParent)
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
	result, err := database.DB.Query(`
		SELECT
		*
		FROM ParentTable
		WHERE IdParent = ?`, IdParent)
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
	result, err := database.DB.Query(`
		SELECT
		*
		FROM ParentTable 
		WHERE IdMember = ?`, idMember)
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

func (parent Parent) ValidateFields(c *fiber.Ctx) error {
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
			return err
		} else {
			continue
		}
	}
	return nil
}

func (parent Parent) GetTotalRows(c *fiber.Ctx) (int, error) {
	var totalRows int
	idMember, err := Member{}.GetIdModel(c)
	if err != nil {
		return 0, err
	}
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE IdMember = ?`, idMember)
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

func (parent Parent) GetAllModels() ([]Parent, error) {
	return nil, nil
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
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM ParentTable 
		WHERE IdParent = ?`, idParent)
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
