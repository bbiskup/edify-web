package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	ssp "github.com/bbiskup/edify/edifact/spec/segment"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var segsTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"SegSpecURL": SegSpecURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	segsTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths(
			"layout.html", "navbar.html",
			"segspecs.html")...,
	))
}

func SegSpecs(w http.ResponseWriter, r *http.Request) {
	segSpecs := make(ssp.SegSpecs, 0, defs.Validator.SegSpecs.Len())
	for _, segSpecId := range defs.Validator.SegSpecs.Ids() {
		segSpec := defs.Validator.SegSpecs.Get(segSpecId)
		segSpecs = append(segSpecs, segSpec)
	}
	sort.Sort(segSpecs)

	err := segsTemplates.ExecuteTemplate(w, "layout", segSpecs)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
