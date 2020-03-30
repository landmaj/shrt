package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"log"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
	tmpl   *template.Template
	db     *sql.DB
}

func (s *server) ListenAndServe() error {
	srv := &http.Server{
		Addr:         "127.0.0.1:5000",
		Handler:      s.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("Application staring...")
	return srv.ListenAndServe()
}

func Run() {
	srv := server{
		router: mux.NewRouter(),
		tmpl:   template.Must(template.ParseGlob("template/*.gohtml")),
		db:     newDatabase(),
	}
	srv.routes()
	log.Fatal(srv.ListenAndServe())
}
