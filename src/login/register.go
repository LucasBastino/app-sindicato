package login

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	"github.com/LucasBastino/app-sindicato/src/errors/customError"
	"github.com/LucasBastino/app-sindicato/src/errors/errorHandler"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	user := c.FormValue("user")
	password := c.FormValue("password")
	// admin, err := strconv.ParseBool(c.FormValue("admin"))
	// if err != nil {
	// 	customError.InternalServerError.Msg = err.Error()
	// 	return errorHandler.HandleError(c, &customError.InternalServerError)
	// }
	writeMember, err := strconv.ParseBool(c.FormValue("write-member"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}
	deleteMember, err := strconv.ParseBool(c.FormValue("delete-member"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}
	writeEnterprise, err := strconv.ParseBool(c.FormValue("write-enterprise"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}
	deleteEnterprise, err := strconv.ParseBool(c.FormValue("delete-enterprise"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}
	writeParent, err := strconv.ParseBool(c.FormValue("write-parent"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}
	deleteParent, err := strconv.ParseBool(c.FormValue("delete-parent"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}
	writePayment, err := strconv.ParseBool(c.FormValue("write-payment"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}
	deletePayment, err := strconv.ParseBool(c.FormValue("delete-payment"))
	if err != nil {
		customError.InternalServerError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.InternalServerError)
	}

	row := database.DB.QueryRow("SELECT IdUser FROM UserTable WHERE Username = ?", user)
	var idUser int
	err = row.Scan(&idUser)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, error:", err)
		return c.Render("register", fiber.Map{"userError": "Nombre de usuario ya existente"})
	}

	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("error generating hash from password")
		// panic(err)
	}

	strHash := string(byteHash)

	insert, err := database.DB.Query(`
	INSERT INTO UserTable 
	(Username, Hash, Admin, WriteMember, DeleteMember, WriteEnterprise, DeleteEnterprise, WriteParent, DeleteParent, WritePayment, DeletePayment)
	VALUES 
	(?, ?, 0, ?, ?, ?, ?, ?, ?, ?, ?)`, user, strHash, writeMember, deleteMember, writeEnterprise, deleteEnterprise, writeParent, deleteParent, writePayment, deletePayment)
	// (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, user, strHash, admin, writeMember, deleteMember, writeEnterprise, deleteEnterprise, writeParent, deleteParent, writePayment, deletePayment)
	if err != nil {
		customError.QueryError.Msg = err.Error()
		return errorHandler.HandleError(c, &customError.QueryError)
	}
	insert.Close()
	return c.Render("login", fiber.Map{})

}
