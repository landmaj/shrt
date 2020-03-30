package app

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
	tmpl   *template.Template
}

func (s *server) ListenAndServe() error {
	srv := &http.Server{
		Addr:         "127.0.0.1:5000",
		Handler:      s.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return srv.ListenAndServe()
}

func Run() {
	srv := server{
		router: mux.NewRouter(),
		tmpl:   template.Must(template.ParseGlob("template/*.gohtml")),
	}
	srv.routes()
	log.Fatal(srv.ListenAndServe())
}
