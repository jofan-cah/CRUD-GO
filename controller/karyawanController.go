package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type Karyawan struct {
	Index   int
	ID      int
	Name    string
	NPWP    string
	Adddres string
}

func NewIndexKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM karyawan")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		defer rows.Close()

		var karyawans []Karyawan
		for rows.Next() {
			var karyawan Karyawan

			err = rows.Scan(
				&karyawan.ID,
				&karyawan.Name,
				&karyawan.NPWP,
				&karyawan.Adddres,
			)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			karyawans = append(karyawans, karyawan)
		}

		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["karyawans"] = karyawans

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func NewCreateKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			adddres := r.Form["adddres"][0]
			_, err := db.Exec("INSERT INTO karyawan(name,npwp,adddres) values(?,?,?)", name, npwp, adddres)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/karyawan", http.StatusMovedPermanently)

		} else if r.Method == "GET" {

			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}

	}
}

func NewUpdateKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid ID"))
				return
			}

			r.ParseForm()
			name := r.Form.Get("name")
			npwp := r.Form.Get("npwp")
			adddres := r.Form.Get("adddres")

			_, err = db.Exec("UPDATE karyawan SET name=?, npwp=?, adddres=? WHERE id=?", name, npwp, adddres, id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			http.Redirect(w, r, "/karyawan", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid ID"))
				return
			}

			row := db.QueryRow("SELECT name, npwp, adddres FROM karyawan WHERE id = ?", id)
			if row.Err() != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(row.Err().Error()))
				return
			}

			var karyawan Karyawan
			err = row.Scan(&karyawan.Name, &karyawan.NPWP, &karyawan.Adddres)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			karyawan.ID = id

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			data := make(map[string]any)
			data["karyawan"] = karyawan

			err = tmpl.Execute(w, data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}
	}
}

func NewDeleteKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		_, err := db.Exec("DELETE FROM karyawan WHERE id = ?", id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/karyawan", http.StatusMovedPermanently)
	}
}
