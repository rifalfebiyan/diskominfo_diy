package controller

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"
)

func NewCreateTamuController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			id_tamu := r.Form["id_tamu"][0]
			nama_tamu := r.Form["nama_tamu"][0]
			jenis_kelamin := r.Form["jenis_kelamin"][0]
			alamat := r.Form["alamat"][0]

			_, err := db.Exec("INSERT INTO tbl_tamu (id_tamu, nama_tamu, jenis_kelamin, alamat) VALUES (?, ?, ?, ?)", id_tamu, nama_tamu, jenis_kelamin, alamat)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			//lokasi yang dituju
			http.Redirect(w, r, "/tamu", http.StatusMovedPermanently)

			// w.Write([]byte("test post"))
			return
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}
