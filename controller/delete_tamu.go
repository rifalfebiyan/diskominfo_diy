package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteTamuController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_tamu := r.URL.Query().Get("id_tamu")
		// r.ParseForm()

		// // id_tamu := r.Form["id_tamu"][0]
		// nama_tamu := r.Form["nama_tamu"][0]
		// jenis_kelamin := r.Form["jenis_kelamin"][0]
		// alamat := r.Form["alamat"][0]

		_, err := db.Exec("DELETE FROM tbl_tamu WHERE id_tamu = ?", id_tamu)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		//lokasi yang dituju
		http.Redirect(w, r, "/tamu", http.StatusMovedPermanently)
		return

	}
}
