package login

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/LucasBastino/app-sindicato/src/database"
	er "github.com/LucasBastino/app-sindicato/src/errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	user := c.FormValue("user")
	password := c.FormValue("password")
	// admin, err := strconv.ParseBool(c.FormValue("admin"))
	// if err != nil {
	// 	er.InternalServerError.Msg = err.Error()
	// 	return er.CheckError(c, er.InternalServerError)
	// }
	writeMember, err := strconv.ParseBool(c.FormValue("write-member"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	deleteMember, err := strconv.ParseBool(c.FormValue("delete-member"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	writeEnterprise, err := strconv.ParseBool(c.FormValue("write-enterprise"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	deleteEnterprise, err := strconv.ParseBool(c.FormValue("delete-enterprise"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	writeParent, err := strconv.ParseBool(c.FormValue("write-parent"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	deleteParent, err := strconv.ParseBool(c.FormValue("delete-parent"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	writePayment, err := strconv.ParseBool(c.FormValue("write-payment"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}
	deletePayment, err := strconv.ParseBool(c.FormValue("delete-payment"))
	if err != nil {
		er.InternalServerError.Msg = err.Error()
		return er.CheckError(c, er.InternalServerError)
	}

	fmt.Println("user:", user)
	fmt.Println("password:", password)
	// fmt.Println("admin:", admin)
	fmt.Println("write member:", writeMember)
	fmt.Println("delete member:", deleteMember)
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
		panic(err)
	}

	strHash := string(byteHash)
	fmt.Println(strHash)

	insert, err := database.DB.Query(`
	INSERT INTO UserTable 
	(Username, Hash, Admin, WriteMember, DeleteMember, WriteEnterprise, DeleteEnterprise, WriteParent, DeleteParent, WritePayment, DeletePayment)
	VALUES 
	(?, ?, 0, ?, ?, ?, ?, ?, ?, ?, ?)`, user, strHash, writeMember, deleteMember, writeEnterprise, deleteEnterprise, writeParent, deleteParent, writePayment, deletePayment)
	// (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, user, strHash, admin, writeMember, deleteMember, writeEnterprise, deleteEnterprise, writeParent, deleteParent, writePayment, deletePayment)
	if err != nil {
		er.QueryError.Msg = err.Error()
		return er.CheckError(c, er.QueryError)
	}
	insert.Close()
	return c.Render("login", fiber.Map{})

}
