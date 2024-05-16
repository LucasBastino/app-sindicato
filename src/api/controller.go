package api

import (
	"database/sql"
	"net/http"
)

type Controller struct {
	db *sql.DB
}

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hola"))
}

func (c *Controller) getUsers(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) insertUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	insert, err := db.Query("CREATE TABLE afiliado(IdAfiliado INT PRIMARY KEY AUTO_INCREMENT,Nombre VARCHAR(45),Edad INT);")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}
