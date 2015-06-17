package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	dsp "github.com/bbiskup/edify/edifact/spec/dataelement"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var compositeDataElemTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"DataElemSpecURL": DataElemSpecURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	compositeDataElemTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths(
			"layout.html", "navbar.html",
			"compositedataelemspecs.html")...,
	))
}

func CompositeDataElemSpecs(w http.ResponseWriter, r *http.Request) {
	sourceSpecs := defs.SpecParser.CompositeDataElemSpecs
	compositeDataElemSpecs := make(dsp.CompositeDataElemSpecs, 0, len(sourceSpecs))
	for _, compositeDataElemSpec := range sourceSpecs {
		compositeDataElemSpecs = append(compositeDataElemSpecs, compositeDataElemSpec)
	}
	sort.Sort(compositeDataElemSpecs)

	err := compositeDataElemTemplates.ExecuteTemplate(w, "layout", compositeDataElemSpecs)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
