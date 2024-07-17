package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Member struct {
	IdMember int
	Name     string
	DNI      string
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

func (m Member) RenderFileTemplate(w http.ResponseWriter, path string) {

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, m)
}

func (m Member) RenderTableTemplate(w http.ResponseWriter, path string, modelList []Member) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		// TmplError{path}.Error(err)
		fmt.Println("error parsing file", path)
	}
	tmpl.Execute(w, modelList)
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

func (m Member) SearchInDB(r *http.Request, DB *sql.DB) []Member {
	searchKey := r.FormValue("search-key")
	var members []Member
	var member Member

	result, err := DB.Query(fmt.Sprintf(`SELECT * FROM MemberTable WHERE Name LIKE '%%%s%%' OR DNI LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching member in DB")
	}
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning member")
		}
		members = append(members, member)
	}
	defer result.Close()
	return members
}

func (m Member) SearchAllModels(DB *sql.DB) []Member {
	member := Member{}
	members := []Member{}
	result, err := DB.Query("SELECT IdMember, Name, DNI FROM MemberTable")
	if err != nil {
		fmt.Println("error searching all members")
	}
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning data from member")
		}
		members = append(members, member)
	}
	return members
}

func (m Member) GetIdModel(r *http.Request) int {
	IdMemberStr := r.PathValue("IdMember")
	IdMember, err := strconv.Atoi(IdMemberStr)
	if err != nil {
		fmt.Println("error converting type")
		panic(err)
	}
	return IdMember
}
