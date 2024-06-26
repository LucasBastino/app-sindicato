package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/LucasBastino/app-sindicato/src/models"
	// "syscall/js"
)

type Controller struct {
	DB *sql.DB
}

func (c *Controller) createMember(w http.ResponseWriter, r *http.Request) {
	newMember := parseMember(r)

	fmt.Println(newMember)

	insert, err := c.DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI) VALUES ('%s','%s')", newMember.Name, newMember.DNI))
	if err != nil {
		fmt.Println("error inserting data in database")
		log.Panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/index", http.StatusSeeOther) // poner un status de redirect (30X), sino no funciona
	// c.renderMemberList(w, r) // esto tambien funciona
}

func (c *Controller) deleteMember(w http.ResponseWriter, r *http.Request) {
	fmt.Println("estamos en deleteMember")
	IdMember := r.PathValue("IdMember")
	fmt.Println(IdMember)
	delete, err := c.DB.Query(fmt.Sprintf("DELETE FROM MemberTable WHERE IdMember = '%s'", IdMember))
	if err != nil {
		fmt.Printf("error deleting member %s from database", IdMember)
		panic(err)
	}
	delete.Close()

	c.renderMemberTable(w, r)
}

func (c *Controller) editMember(w http.ResponseWriter, r *http.Request) {
	memberEdited := parseMember(r)
	IdMember := r.PathValue("IdMember")
	update, err := c.DB.Query(fmt.Sprintf("UPDATE MemberTable SET Name = '%s', DNI = '%s' WHERE IdMember = '%s'", memberEdited.Name, memberEdited.DNI, IdMember))
	if err != nil {
		fmt.Println("error updating member", memberEdited.Name)
		panic(err)
	}
	update.Close()
	// no puedo hacer esto ↓ porque estoy en POST, no puedo redireccionar
	http.Redirect(w, r, "/index", http.StatusSeeOther) // con este status me anda, con otros de 300 no
}

func (c *Controller) editParent(w http.ResponseWriter, r *http.Request) {
	IdParent := r.PathValue("IdParent")
	Name := r.FormValue("name")
	Rel := r.FormValue("rel")

	update, err := c.DB.Query(fmt.Sprintf("UPDATE ParentTable SET Name = '%s', Rel = '%s' WHERE IdParent = '%s'", Name, Rel, IdParent))
	if err != nil {
		fmt.Println("error updating parent")
		panic(err)
	}
	update.Close()

	c.renderParentFile(w, r)
}

func (c *Controller) createEnterprise(w http.ResponseWriter, r *http.Request) {
	var enterprise models.Enterprise
	enterprise.Name = r.FormValue("name")
	enterprise.Address = r.FormValue("address")
	// parseEnterprise()

	insert, err := c.DB.Query(fmt.Sprintf("INSERT INTO EnterpriseTable (Name, Address) VALUES ('%s', '%s')", enterprise.Name, enterprise.Address))
	if err != nil {
		fmt.Println("error inserting data to database")
		panic(err)
	}
	defer insert.Close()

	tmpl, err := template.ParseFiles("src/views/files/enterpriseFile.html")
	if err != nil {
		fmt.Println("error parsing file enterpriseFile.html")
		panic(err)
	}
	tmpl.Execute(w, enterprise)

}
