package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var simpleDataElemSpecTemplates *template.Template

func init() {
	funcMap := template.FuncMap{}
	t := template.New("layout.html").Funcs(funcMap)
	simpleDataElemSpecTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "simpledataelemspec.html")...,
	))
}

func SimpleDataElemSpec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataElemSpecID := vars["id"]

	data := map[string]interface{}{
		"dataElemSpec": defs.SpecParser.SimpleDataElemSpecs[dataElemSpecID],
	}

	err := simpleDataElemSpecTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
