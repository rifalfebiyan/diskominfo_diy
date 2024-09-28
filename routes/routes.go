package routes

import (
	"database/sql"
	"net/http"

	"github.com/rifalfebiyan/diskominfo_diy/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/tamu", controller.NewIndexTamu(db))
	server.HandleFunc("/tamu/create", controller.NewCreateTamuController(db))
	server.HandleFunc("/tamu/update", controller.NewUpdateTamuController(db))
	server.HandleFunc("/tamu/delete", controller.NewDeleteTamuController(db))

}
