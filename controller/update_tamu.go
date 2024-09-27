package controller

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"
)

func NewUpdateTamuController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id_tamu := r.URL.Query().Get("id_tamu")
			r.ParseForm()

			// id_tamu := r.Form["id_tamu"][0]
			nama_tamu := r.Form["nama_tamu"][0]
			jenis_kelamin := r.Form["jenis_kelamin"][0]
			alamat := r.Form["alamat"][0]

			_, err := db.Exec("UPDATE tbl_tamu SET nama_tamu = ?, jenis_kelamin = ?, alamat = ? WHERE id_tamu = ?", nama_tamu, jenis_kelamin, alamat, id_tamu)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			//lokasi yang dituju
			http.Redirect(w, r, "/tamu", http.StatusMovedPermanently)
			return

		} else if r.Method == "GET" {
			//mengarahkan ID Tamunya
			id_tamu := r.URL.Query().Get("id_tamu")
			row := db.QueryRow("SELECT nama_tamu, jenis_kelamin, alamat FROM tbl_tamu WHERE id_tamu = ?", id_tamu)
			// if row.Err != nil {
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	w.Write([]byte(row.Err().Error()))
			// 	return
			// }

			var tamu Tamu
			err := row.Scan(

				&tamu.NamaTamu,
				&tamu.JenisKelamin,
				&tamu.Alamat,
			)

			tamu.IdTamu = id_tamu
			// if row.Err != nil {
			// 	w.WriteHeader(http.StatusInternalServerError)
			// 	w.Write([]byte(row.Err().Error()))
			// 	return
			// }
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			//bikin ronemap
			data := make(map[string]any)
			data["tamu"] = tamu

			err = tmpl.Execute(w, data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}
