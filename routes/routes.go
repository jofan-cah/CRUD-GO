package routes

import (
	"database/sql"
	"net/http"

	"github.com/jofan-cah/crud-go-kar/controller"
)

func MapKaryawanRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/karyawan", controller.NewIndexKaryawan(db))
	server.HandleFunc("/karyawan/create", controller.NewCreateKaryawan(db))
	server.HandleFunc("/karyawan/update", controller.NewUpdateKaryawan(db))
	server.HandleFunc("/karyawan/delete", controller.NewDeleteKaryawan(db))
}
