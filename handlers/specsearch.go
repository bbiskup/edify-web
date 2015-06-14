package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"html/template"
	"log"
	"net/http"
)

var specSearchTemplates *template.Template

func init() {
	specSearchTemplates = template.Must(template.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "specsearch.html")...,
	))
}

func SpecSearch(w http.ResponseWriter, r *http.Request) {
	err := specSearchTemplates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
