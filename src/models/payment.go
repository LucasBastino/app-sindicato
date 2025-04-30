package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	"github.com/gofiber/fiber/v2"
)

type Payment struct {
	IdPayment    int
	Month        string
	Year         string
	Status       string
	Amount       int
	PaymentDate  string
	Commentary   string
	IdEnterprise int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (payment Payment) InsertModel() (Payment, error) {
	insert, err := database.DB.Query(`
		INSERT INTO PaymentTable 
		(Month,
		Year,
		Status, 
		Amount, 
		PaymentDate, 
		Commentary,
		IdEnterprise)
		VALUES (?,?,?,?,?, ?, ?)`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Commentary,
		payment.IdEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Payment{}, er.QueryError
	}
	insert.Close()
	result, err := database.DB.Query(`
		SELECT *
		FROM PaymentTable
		WHERE IdPayment = (SELECT LAST_INSERT_ID())`)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Payment{}, er.QueryError
	}
	p, _, err := payment.ScanResult(result, true)
	if err != nil {
		return Payment{}, err
	}
	return p, nil
}

func (payment Payment) DeleteModel() error {
	delete, err := database.DB.Query(`
		DELETE FROM PaymentTable
		WHERE IdPayment = ?`, payment.IdPayment)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return er.QueryError
	}
	defer delete.Close()
	return nil
}

func (payment Payment) UpdateModel() (Payment, error) {
	update, err := database.DB.Query(`
		UPDATE PaymentTable 
		SET 
		Month = ?, 
		Year = ?,
		Status = ?, 
		Amount = ?, 
		PaymentDate = ?, 
		Commentary = ? 
		WHERE IdPayment = ?`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Commentary,
		payment.IdPayment)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Payment{}, er.QueryError
	}
	update.Close()
	result, err := database.DB.Query("SELECT * FROM PaymentTable WHERE IdPayment = ?", payment.IdPayment)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Payment{}, er.QueryError
	}
	p, _, err := payment.ScanResult(result, true)
	if err != nil {
		return Payment{}, err
	}
	return p, nil
}

func (payment Payment) GetIdModel(c *fiber.Ctx) (int, error) {
	params := struct {
		IdPayment int `params:"IdPayment"`
	}{}

	err := c.ParamsParser(&params)
	if err != nil {
		er.ParamsError.Msg = err.Error()
		return 0, er.ParamsError
	}

	return params.IdPayment, nil
}

func (payment Payment) SearchOneModelById(c *fiber.Ctx) (Payment, error) {
	IdPayment, err := payment.GetIdModel(c)
	if err != nil {
		return Payment{}, err
	}
	result, err := database.DB.Query(`
		SELECT
		*
		FROM PaymentTable
		WHERE IdPayment = ?`, IdPayment)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Payment{}, er.QueryError
	}

	p, _, err := payment.ScanResult(result, true)
	if err != nil {
		return Payment{}, err
	}
	return p, nil

}

func (payment Payment) SearchModels(c *fiber.Ctx, year int) ([]Payment, string, error) {
	IdEnterprise, err := Enterprise{}.GetIdModel(c)
	if err != nil {
		return nil, "", err
	}
	yearStr := strconv.Itoa(year)
	result, err := database.DB.Query(`
	SELECT
	*
	FROM PaymentTable 
	WHERE Year = ? AND IdEnterprise = ?
	ORDER BY Month ASC
	`, yearStr, IdEnterprise)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return nil, "", er.QueryError
	}
	_, pp, err := payment.ScanResult(result, false)
	if err != nil {
		return nil, "", err
	}
	return pp, "", nil
}

func (payment Payment) ValidateFields(c *fiber.Ctx) error {
	validateFunctions := []func(*fiber.Ctx) error{
		validatePayment,
		validateStatus,
		validatePaymentAmount,
		validatePaymentDate,
		validateCommentary,
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

func (payment Payment) GetTotalRows(c *fiber.Ctx) (int, error) {
	return 0, nil
}

func (payment Payment) GetFiberMap(Payments []Payment, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return nil
}

func (payment Payment) GetAllModels() ([]Payment, error) {
	return nil, nil
}

func (payment Payment) ScanResult(result *sql.Rows, onlyOne bool) (Payment, []Payment, error) {
	var p Payment
	var pp []Payment
	for result.Next() {
		err := result.Scan(
			&p.IdPayment,
			&p.Month,
			&p.Year,
			&p.Status,
			&p.Amount,
			&p.PaymentDate,
			&p.Commentary,
			&p.IdEnterprise,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			er.ScanError.Msg = err.Error()
			return Payment{}, nil, er.ScanError
		}
		if !onlyOne {
			pp = append(pp, p)
		}
	}
	result.Close()
	return p, pp, nil
}

func (payment Payment) CheckDeleted(idPayment int) (bool, error) {
	var totalRows int
	// row := database.DB.QueryRow(fmt.Sprintf(`
	// 	SELECT COUNT(*) FROM PaymentTable
	// 	WHERE IdPayment = '%d'`, p.IdPayment))
	row := database.DB.QueryRow(`
		SELECT COUNT(*) FROM PaymentTable 
		WHERE IdPayment = ?`, idPayment)
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
