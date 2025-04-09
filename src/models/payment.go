package models

import (
	"database/sql"
	"fmt"
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
	insert, err := database.DB.Query(fmt.Sprintf(`
		INSERT INTO PaymentTable 
		(Month,
		Year,
		Status, 
		Amount, 
		PaymentDate, 
		Commentary,
		IdEnterprise)
		VALUES ('%s','%s','%s','%d','%s', '%s', '%d')`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Commentary,
		payment.IdEnterprise))
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
	delete, err := database.DB.Query(fmt.Sprintf(`
		DELETE FROM PaymentTable
		WHERE IdPayment = '%d'`, payment.IdPayment))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return er.QueryError
	}
	defer delete.Close()
	return nil
}

func (payment Payment) UpdateModel() (Payment, error) {
	update, err := database.DB.Query(fmt.Sprintf(`
		UPDATE PaymentTable 
		SET 
		Month = '%s', 
		Year = '%s',
		Status = '%s', 
		Amount = '%d', 
		PaymentDate = '%s', 
		Commentary = '%s' 
		WHERE IdPayment = '%d'`,
		payment.Month,
		payment.Year,
		payment.Status,
		payment.Amount,
		payment.PaymentDate,
		payment.Commentary,
		payment.IdPayment))
	if err != nil {
		er.QueryError.Msg = err.Error()
		return Payment{}, er.QueryError
	}
	update.Close()
	result, err := database.DB.Query(fmt.Sprintf("SELECT * FROM PaymentTable WHERE IdPayment = %d", payment.IdPayment))
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
	result, err := database.DB.Query(fmt.Sprintf(`
		SELECT
		*
		FROM PaymentTable
		WHERE IdPayment = '%d'`, IdPayment))
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
	result, err := database.DB.Query(fmt.Sprintf(`
	SELECT
	*
	FROM PaymentTable 
	WHERE Year = '%s' AND IdEnterprise = '%d'
	ORDER BY Month ASC
	`, yearStr, IdEnterprise))
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

func (payment Payment) ValidateFields(c *fiber.Ctx) (map[string]string, error) {
	errorMap := map[string]string{}
	var valid bool
	var err string

	if valid, err = validatePayment(c); !valid {
		errorMap["payment"] = err
	}
	if valid, err = validateStatus(c); !valid {
		errorMap["status"] = err
	}
	if valid, err = validatePaymentAmount(c); !valid {
		errorMap["amount"] = err
	}
	if valid, err = validatePaymentDate(c); !valid {
		errorMap["paymentDate"] = err
	}
	if valid, err = validateCommentary(c); !valid {
		errorMap["commentary"] = err
	}
	if len(errorMap) > 1 {

		return errorMap, er.ValidationError
	}
	return errorMap, nil
}

func (payment Payment) GetTotalRows(c *fiber.Ctx) int {
	return 0
}

func (payment Payment) GetFiberMap(Payments []Payment, searchKey string, currentPage, someBefore, someAfter, totalPages int, totalPagesArray []int) fiber.Map {
	return nil
}

func (payment Payment) GetAllModels() []Payment {
	return nil
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
	row := database.DB.QueryRow(fmt.Sprintf(`
		SELECT COUNT(*) FROM PaymentTable 
		WHERE IdPayment = '%d'`, idPayment))
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
