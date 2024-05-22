package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/models"
)

type Controller struct {
	DB *sql.DB
}

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hola"))
}

func (c *Controller) getUsers(w http.ResponseWriter, r *http.Request) {

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
	update, err := c.DB.Query("UPDATE afiliado SET Nombre = '%v', Edad = '%v' WHERE IdAfiliado = '%v' ", afiliado.Nombre, afiliado.Edad, id)
	if err != nil {
		fmt.Println("error updating afiliado")
		panic(err)
	}
	update.Close()
	fmt.Println(afiliado)
}

func (c *Controller) createTable() {

	insert, err := c.DB.Query("CREATE TABLE afiliado(IdAfiliado INT PRIMARY KEY AUTO_INCREMENT,Nombre VARCHAR(45),Edad INT);")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}
