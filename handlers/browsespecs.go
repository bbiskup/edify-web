package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"html/template"
	"log"
	"net/http"
)

var browsSpecsTemplates *template.Template

func init() {
	browsSpecsTemplates = template.Must(template.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "browsespecs.html")...,
	))
}

func BrowseSpecs(w http.ResponseWriter, r *http.Request) {
	err := browsSpecsTemplates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
