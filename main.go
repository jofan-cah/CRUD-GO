package main

import (
	"net/http"

	"github.com/jofan-cah/crud-go-kar/database"
	"github.com/jofan-cah/crud-go-kar/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapKaryawanRoutes(server, db)

	http.ListenAndServe(":8080", server)

}
