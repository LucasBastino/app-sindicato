package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
)

type Member struct {
	IdMember int
	Name     string
	DNI      string
}

type Parent struct {
	IdParent int
	Name     string
	Rel      string
	IdMember int
}

type Enterprise struct {
	IdEnterprise int
	Name         string
	Address      string
}

type DBError struct {
	Statement string
	Model     string
}

func (m Member) Imprimir() {
	fmt.Println(m)
}

func (newMember Member) InsertInDB(DB *sql.DB) {
	insert, err := DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI) VALUES ('%s','%s')", newMember.Name, newMember.DNI))
	if err != nil {
		// DBError{"INSERT MEMBER"}.Error(err)
		fmt.Println("error insertando en la DB")
	}
	defer insert.Close()
}

func (m Member) RenderTemplate(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, m)
}

func (m Member) DeleteFromDB(DB *sql.DB) {
	delete, err := DB.Query(fmt.Sprintf("DELETE FROM MemberTable WHERE IdMember = '%v'", m.IdMember))
	if err != nil {
		// DBError{"DELETE MEMBER"}.Error(err)
		fmt.Println("error deleting member")
	}
	defer delete.Close()

}

func (m Member) UpdateInDB(IdMember int, DB *sql.DB) {
	update, err := DB.Query(fmt.Sprintf("UPDATE MemberTable SET Name = '%s', DNI = '%s' WHERE IdMember = '%v'", m.Name, m.DNI, IdMember))
	if err != nil {
		// DBError{"UPDATE MEMBER"}.Error(err)
		fmt.Println("error updating member")
		panic(err)
	}
	update.Close()
}
