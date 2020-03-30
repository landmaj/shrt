package app

import (
	"net/http"
)

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
	return func(w http.ResponseWriter, r *http.Request) {
		link := r.FormValue("link")
		id, err := generateId(link)
		if err != nil {
			s.tmpl.ExecuteTemplate(w, "index.gohtml", data{
				Link:    link,
				Enabled: true,
				Error:   err,
			})
		} else {
			s.tmpl.ExecuteTemplate(w, "index.gohtml", data{
				Link: id,
			})
		}
	}
}
