package handlers

import (
	"fmt"
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
	funcMap := template.FuncMap{"msgSpecURL": msgSpecURL}
	t := template.New("layout.html").Funcs(funcMap)
	specSearchTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "specsearch.html")...,
	))
}

func msgSpecURL(msgSpec *msp.MsgSpec) string {
	return fmt.Sprintf("/specs/message/%s", msgSpec.Id)
}

// Search term in message specifications
func searchMsgSpecs(w http.ResponseWriter, searchTerm string) []*msp.MsgSpec {
	result := msp.MsgSpecs{}
	for msgId, msgSpec := range defs.Validator.MsgSpecs {
		if strings.Contains(strings.ToLower(msgId), searchTerm) || strings.Contains(strings.ToLower(msgSpec.Name), searchTerm) {
			result = append(result, msgSpec)
		}
	}
	sort.Sort(result)
	return result
}

func SpecSearch(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	searchTerm := strings.ToLower(r.FormValue("searchterm"))
	data := map[string]interface{}{}
	msgSpecs := searchMsgSpecs(w, searchTerm)
	log.Printf("Found %d message specs for search term %s", len(msgSpecs), searchTerm)
	data["msgSpecs"] = msgSpecs

	err := specSearchTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
