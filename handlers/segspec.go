package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var segSpecTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"DataElemSpecURL": DataElemSpecURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	segSpecTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "segspec.html")...,
	))
}

func SegSpec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	segSpecID := vars["id"]

	data := map[string]interface{}{
		"segSpec": defs.Validator.SegSpecs.Get(segSpecID),
	}

	err := segSpecTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
