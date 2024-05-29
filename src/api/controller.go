package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "syscall/js"
	"text/template"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type Controller struct {
	DB *sql.DB
}

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/views/home.html"))
	tmpl.Execute(w, nil)
}

func (c *Controller) getUsers(w http.ResponseWriter, r *http.Request) {
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

func createTemplate(path string) *template.Template {
	tmpl, err := template.New("newTemplate").ParseFiles(path)
	if err != nil {
		fmt.Println("error creating template from", path)
		log.Panic(err.Error())
	}
	return tmpl
}

func execTemplate(w http.ResponseWriter, data any, tmpl *template.Template, file string) {
	// w.Header().Set("Content-Type", "application/json")
	err := tmpl.ExecuteTemplate(w, file, data)
	if err != nil {
		fmt.Println("error executing template")
		log.Panic(err)
	}
}
func ImprimirAlgo() func() {
	return func() {
		fmt.Println("holaaaaaa")
	}
}

func (c *Controller) test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Println("testttttttttttttt")
}

func (c *Controller) insertUser(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	age := r.PathValue("age")
	insert, err := c.DB.Query(fmt.Sprintf("INSERT INTO member (name, age) VALUES ('%s', '%s')", name, age))
	if err != nil {
		fmt.Println("error inserting data")
		panic(err.Error())
	}
	insert.Close()
}

func (c *Controller) updateUser(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	// var member models.Member
	// json.NewDecoder(r.Body).Decode(&member)
	// update, err := c.DB.Query(fmt.Sprintf("UPDATE member SET name = '%v', age = '%v' WHERE idMember = '%v' ", member.Name, member.DNI, id))
	// if err != nil {
	// 	fmt.Println("error updating Member")
	// 	panic(err)
	// }
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print("error reading body")
		log.Panic(err)
	}
	fmt.Println(body)
	fmt.Println(string(body))
	// update.Close()
}

func (c *Controller) createTable(w http.ResponseWriter, r *http.Request) {

	insert, err := c.DB.Query("CREATE TABLE member(idMember INT PRIMARY KEY AUTO_INCREMENT,name VARCHAR(45),age INT);")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

func (c *Controller) renderCreateMemberForm(w http.ResponseWriter, req *http.Request) {
	tmpl := createTemplate("src/views/createMemberForm.html")
	execTemplate(w, nil, tmpl, "createMemberForm.html")
}

// type MemberRequest struct {
// 	IdMember string `json:"idmember"`
// 	Name     string `json:"name"`
// 	DNI      string `json:"dni"`
// }

func (c *Controller) createMember(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// name := req.FormValue("Name")
	// dni := req.FormValue("DNI")
	// newMember := models.Member{Name: name, DNI: dni}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error reading body")
		log.Panic(err.Error())
	}
	fmt.Println(string(body))

	var newMember models.Member

	err = json.Unmarshal(body, &newMember)
	if err != nil {
		fmt.Println("error unmarshalling body")
		panic(err)
	}

	// err = json.NewDecoder(req.Body).Decode(&newMember)
	// if err != nil {
	// 	fmt.Println("error decoding member")
	// 	log.Panic(err)
	// }
	fmt.Println(newMember)
	// w.Write([]byte(req.Body))
	// insert, err := c.DB.Query(fmt.Sprintf("INSERT INTO MemberTable (Name, DNI) VALUES ('%s', '%s')", newMember.Name, newMember.DNI))
	// if err != nil {
	// 	fmt.Println("error inserting data in database")
	// 	log.Panic(err.Error())
	// }
	// defer insert.Close()
}

// funcmap := map[string]interface{}{
// 	"Imprimir": ImprimirAlgo,
// }
// tmpl, err := template.New("tmplUsers").Funcs(funcmap).ParseFiles("src/views/index.html")
// if err != nil {
// 	fmt.Println("error creating template")
// 	log.Panicln(err.Error())
// }
