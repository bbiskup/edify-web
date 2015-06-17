package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	dsp "github.com/bbiskup/edify/edifact/spec/dataelement"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var simpleDataElemTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"DataElemSpecURL": DataElemSpecURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	simpleDataElemTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths(
			"layout.html", "navbar.html",
			"simpledataelemspecs.html")...,
	))
}

func SimpleDataElemSpecs(w http.ResponseWriter, r *http.Request) {
	sourceSpecs := defs.SpecParser.SimpleDataElemSpecs
	simpleDataElemSpecs := make(dsp.SimpleDataElemSpecs, 0, len(sourceSpecs))
	for _, simpleDataElemSpec := range sourceSpecs {
		simpleDataElemSpecs = append(simpleDataElemSpecs, simpleDataElemSpec)
	}
	sort.Sort(simpleDataElemSpecs)

	err := simpleDataElemTemplates.ExecuteTemplate(w, "layout", simpleDataElemSpecs)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
