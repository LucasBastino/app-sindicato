package main

import (
	"fmt"
	"net/http"

	"github.com/LucasBastino/app-sindicato/src/api"
	"github.com/LucasBastino/app-sindicato/src/connection"
)

func main() {
	r := http.NewServeMux()
	svr := &http.Server{
		Addr:    ":8085",
		Handler: r,
	}
	c := api.Controller{}
	c.DB = connection.CreateConnection()

	c.RegisterRoutes(r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listen on port 8085")
	svr.ListenAndServe()
}
