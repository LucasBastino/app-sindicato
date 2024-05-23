package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
	var afiliado models.Afiliado
	var afiliados []models.Afiliado
	result, err := c.DB.Query("SELECT Nombre, Edad FROM afiliado ")
	if err != nil {
		fmt.Println("error getting users")
	}
	for result.Next() {
		err = result.Scan(&afiliado.Nombre, &afiliado.Edad)
		if err != nil {
			fmt.Println("error scanning result")
			panic(err.Error())
		}
		afiliados = append(afiliados, afiliado)
	}
	tmpl := template.Must(template.ParseFiles("src/views/index.html"))
	tmpl.Execute(w, afiliados)
}

func (c *Controller) insertUser(w http.ResponseWriter, r *http.Request) {
	nombre := r.PathValue("nombre")
	edad := r.PathValue("edad")
	insert, err := c.DB.Query(fmt.Sprintf("INSERT INTO afiliado (Nombre, Edad) VALUES ('%s', '%s')", nombre, edad))
	if err != nil {
		fmt.Println("error inserting data")
		panic(err.Error())
	}
	insert.Close()
}

func (c *Controller) updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var afiliado models.Afiliado
	json.NewDecoder(r.Body).Decode(&afiliado)
	update, err := c.DB.Query(fmt.Sprintf("UPDATE afiliado SET Nombre = '%v', Edad = '%v' WHERE IdAfiliado = '%v' ", afiliado.Nombre, afiliado.Edad, id))
	if err != nil {
		fmt.Println("error updating afiliado")
		panic(err)
	}
	update.Close()
}

func (c *Controller) createTable() {

	insert, err := c.DB.Query("CREATE TABLE afiliado(IdAfiliado INT PRIMARY KEY AUTO_INCREMENT,Nombre VARCHAR(45),Edad INT);")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}
