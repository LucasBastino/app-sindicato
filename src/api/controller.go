package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

	c.renderMemberList(w, r)
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
