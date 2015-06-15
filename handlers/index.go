package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"html/template"
	"log"
	"net/http"
)

var indexTemplates *template.Template

func init() {
	indexTemplates = template.Must(template.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "index.html")...,
	))
}

func Index(w http.ResponseWriter, r *http.Request) {
	err := indexTemplates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
