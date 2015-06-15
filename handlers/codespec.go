package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var codeSpecTemplates *template.Template

func init() {
	funcMap := template.FuncMap{}
	t := template.New("layout.html").Funcs(funcMap)
	codeSpecTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "codespec.html")...,
	))
}

func CodeSpec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	codeSpecID := vars["id"]

	data := map[string]interface{}{
		"codeSpec": defs.SpecParser.CodesSpecs[codeSpecID],
	}

	err := codeSpecTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
