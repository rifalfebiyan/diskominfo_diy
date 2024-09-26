package main

import (
	"net/http"

	"github.com/rifalfebiyan/diskominfo_diy/database"
	"github.com/rifalfebiyan/diskominfo_diy/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)
	http.ListenAndServe(":8080", server)
}
