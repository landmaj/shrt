package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("template/*.gohtml"))

func indexGet(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func indexPost(w http.ResponseWriter, r *http.Request) {
	link := r.FormValue("link")
	tmpl.ExecuteTemplate(w, "index.gohtml", link)
}

func shortcut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}
