package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("template/*.gohtml"))

type data struct {
	Link    string
	Enabled bool
	Error   error
}

func indexGet(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", data{
		Link:    r.FormValue("link"),
		Enabled: true,
	})
}

func indexPost(w http.ResponseWriter, r *http.Request) {
	link := r.FormValue("link")
	id, err := generateId(link)
	if err != nil {
		tmpl.ExecuteTemplate(w, "index.gohtml", data{
			Link:    link,
			Enabled: true,
			Error:   err,
		})
	} else {
		tmpl.ExecuteTemplate(w, "index.gohtml", data{
			Link: id,
		})
	}
}

func shortcut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["shrt"])
}
