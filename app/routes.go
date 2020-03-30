package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const rootDomainEnv = "ROOT_DOMAIN"

type data struct {
	Link    string
	Enabled bool
	Error   error
}

func (s *server) routes() {
	fs := http.FileServer(http.Dir("static/"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	s.router.HandleFunc("/", s.indexGet()).Methods(http.MethodGet)
	s.router.HandleFunc("/", s.indexPost()).Methods(http.MethodPost)
	s.router.HandleFunc("/{shrt}", s.redirectShrt()).Methods(http.MethodGet)
}

func (s *server) indexGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.tmpl.ExecuteTemplate(w, "index.gohtml", data{
			Link:    r.FormValue("link"),
			Enabled: true,
		})
	}
}

func (s *server) indexPost() http.HandlerFunc {
	rootDomain, exists := os.LookupEnv(rootDomainEnv)
	if !exists {
		log.Fatalln("missing environmental variable:", rootDomainEnv)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		link := r.FormValue("link")
		id, err := generateShrt(s.db, link)
		if err != nil {
			s.tmpl.ExecuteTemplate(w, "index.gohtml", data{
				Link:    link,
				Enabled: true,
				Error:   err,
			})
		} else {
			s.tmpl.ExecuteTemplate(w, "index.gohtml", data{
				Link: rootDomain + "/" + id,
			})
		}
	}
}

func (s *server) redirectShrt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		row, err := queryByShrt(s.db, vars["shrt"])
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		} else if err != nil {
			w.WriteHeader(500)
			return
		}
		http.Redirect(w, r, row.url, http.StatusSeeOther)
	}
}
