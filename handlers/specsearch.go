package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	msp "github.com/bbiskup/edify/edifact/spec/message"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

var specSearchTemplates *template.Template

func init() {
	specSearchTemplates = template.Must(template.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "specsearch.html")...,
	))
}

// Search term in message specifications
func searchMsgSpecs(w http.ResponseWriter, searchTerm string) []*msp.MsgSpec {
	result := msp.MsgSpecs{}
	for msgId, msgSpec := range defs.Validator.MsgSpecs {
		if strings.Contains(msgId, searchTerm) || strings.Contains(msgSpec.Name, searchTerm) {
			result = append(result, msgSpec)
		}
	}
	sort.Sort(result)
	return result
}

func SpecSearch(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	searchTerm := r.FormValue("searchterm")
	data := map[string]interface{}{}
	if searchTerm != "" {
		msgSpecs := searchMsgSpecs(w, searchTerm)
		log.Printf("Found %d message specs for search term %s", len(msgSpecs), searchTerm)
		data["msgSpecs"] = msgSpecs
	}

	err := specSearchTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
