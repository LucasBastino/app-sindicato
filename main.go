package main

import (
	"fmt"
	"net/http"

	api "github.com/LucasBastino/app-sindicato/src/api/controller"
	"github.com/LucasBastino/app-sindicato/src/connection"
)

func main() {
	muxer := http.NewServeMux()
	svr := &http.Server{
		Addr:    ":8085",
		Handler: muxer,
	}
	c := api.Controller{}
	c.DB = connection.CreateConnection()

	// buscar como obtener el current PATH, y cambiarle las \ por / o eso de que depende el OS lo cambia solo
	fileServerStatic := http.FileServer(http.Dir("src/static")) // con o sin barra al final es lo mismo
	// si o si es muxer.Handle, con http.Handle no funciona
	muxer.Handle("/static/", http.StripPrefix("/static/", fileServerStatic))
	c.RegisterRoutes(muxer)

	fmt.Println("Listen on port 8085")

	svr.ListenAndServe()
}
