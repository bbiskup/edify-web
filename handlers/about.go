package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"html/template"
	"log"
	"net/http"
)

var aboutTemplates *template.Template

func init() {
	aboutTemplates = template.Must(template.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "about.html")...,
	))
}

func About(w http.ResponseWriter, r *http.Request) {
	err := aboutTemplates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
