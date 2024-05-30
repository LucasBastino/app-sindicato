package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// "syscall/js"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type Controller struct {
	DB *sql.DB
}

func (c *Controller) getMembers(w http.ResponseWriter, r *http.Request) {
	var member models.Member
	var members []models.Member
	result, err := c.DB.Query("SELECT name, age FROM member ")
	if err != nil {
		fmt.Println("error getting users")
	}
	for result.Next() {
		err = result.Scan(&member.Name, &member.DNI)
		if err != nil {
			fmt.Println("error scanning result")
			panic(err.Error())
		}
		members = append(members, member)
	}

	tmpl := createTemplate("src/views/index.html")
	execTemplate(w, members, tmpl, "index.html")

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

	// tmpl, err := template.New("newTemplate").ParseFiles("src/views/memberList.html")
	// if err != nil {
	// 	fmt.Println("error creating template")
	// 	log.Panic(err)
	// }
	http.Redirect(w, r, "/members", http.StatusOK)
}
