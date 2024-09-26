package main

import (
	"net/http"

	"github.com/rifalfebiyan/diskominfo_diy/database"
	"github.com/rifalfebiyan/diskominfo_diy/routes"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server)
	http.ListenAndServe(":8080", server)
}
