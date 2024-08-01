package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Member struct {
	IdMember int
	Name     string
	DNI      string
}

func (m Member) Imprimir() {
	fmt.Println(m)
}

func (newMember Member) InsertModel(DB *sql.DB) {
	insert, err := DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI) VALUES ('%s','%s')", newMember.Name, newMember.DNI))
	if err != nil {
		// DBError{"INSERT MEMBER"}.Error(err)
		fmt.Println("error insertando en la DB")
		panic(err)
	}
	defer insert.Close()
}

func (m Member) DeleteModel(DB *sql.DB) {
	delete, err := DB.Query(fmt.Sprintf("DELETE FROM MemberTable WHERE IdMember = '%v'", m.IdMember))
	if err != nil {
		// DBError{"DELETE MEMBER"}.Error(err)
		fmt.Println("error deleting member")
		panic(err)
	}
	defer delete.Close()

}

func (m Member) EditModel(DB *sql.DB) {
	update, err := DB.Query(fmt.Sprintf("UPDATE MemberTable SET Name = '%s', DNI = '%s' WHERE IdMember = '%v'", m.Name, m.DNI, m.IdMember))
	if err != nil {
		// DBError{"UPDATE MEMBER"}.Error(err)
		fmt.Println("error updating member")
		panic(err)
	}
	defer update.Close()
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

func (m Member) SearchOneModelById(r *http.Request, DB *sql.DB) Member {
	IdMember := m.GetIdModel(r)
	result, err := DB.Query(fmt.Sprintf("SELECT IdMember, Name, DNI FROM MemberTable WHERE IdMember = '%v'", IdMember))
	if err != nil {
		fmt.Println("error searching member by Id")
		panic(err)
	}

	var member Member
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning member")
			panic(err)
		}
	}
	defer result.Close()
	return member
}

func (m Member) SearchModels(r *http.Request, DB *sql.DB) []Member {
	searchKey := r.FormValue("search-key")
	var members []Member
	var member Member

	result, err := DB.Query(fmt.Sprintf(`SELECT * FROM MemberTable WHERE Name LIKE '%%%s%%' OR DNI LIKE '%%%s%%'`, searchKey, searchKey))
	if err != nil {
		fmt.Println("error searching member in DB")
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&member.IdMember, &member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning member")
			panic(err)
		}
		members = append(members, member)
	}
	defer result.Close()
	return members
}

func (m Member) ValidateFields(r *http.Request) map[string]string {
	errorMap := map[string]string{}

	if strings.TrimSpace(r.FormValue("name")) == "" {
		errorMap["name"] = "el campo Nombre no puede estar vacio"
	}
	// consultar que sea alfanumerico
	if strings.TrimSpace(r.FormValue("dni")) == "" {
		errorMap["dni"] = "el campo DNI no puede estar vacio"
	}
	if utf8.RuneCountInString(r.FormValue("dni")) > 8 {
		errorMap["dni"] = "el DNI no puede tener mas de 8 caracteres"
	}
	return errorMap
}

func (m Member) CreateTemplateData(member Member, members []Member, path string, errorMap map[string]string) TemplateData {
	templateData := TemplateData{}
	templateData.Member = member
	templateData.Members = members
	templateData.Path = path
	templateData.ErrorMap = errorMap
	return templateData
}
