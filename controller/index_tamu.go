package controller

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"
)

// deklarasi
type Tamu struct {
	IdTamu       int
	NamaTamu     string
	JenisKelamin string
	Alamat       string
}

func NewIndexTamu(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query untuk menampilkan data dari database, Exec itu untuk insert, delete
		row, err := db.Query("SELECT id_tamu, nama_tamu, jenis_kelamin, alamat FROM tbl_tamu")
		// CEK ERROR
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// defer rows.Close()

		var tamuu []Tamu
		for row.Next() {
			var tamu Tamu

			err := row.Scan(
				&tamu.IdTamu,
				&tamu.NamaTamu,
				&tamu.JenisKelamin,
				&tamu.Alamat,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tamuu = append(tamuu, tamu)

		}

		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["tamuu"] = tamuu

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
