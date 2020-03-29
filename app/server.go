package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Run() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", indexGet).Methods(http.MethodGet)
	r.HandleFunc("/", indexPost).Methods(http.MethodPost)
	r.HandleFunc("/{id}", shortcut).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         "127.0.0.1:5000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
