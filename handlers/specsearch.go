package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	dsp "github.com/bbiskup/edify/edifact/spec/dataelement"
	msp "github.com/bbiskup/edify/edifact/spec/message"
	ssp "github.com/bbiskup/edify/edifact/spec/segment"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

var specSearchTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"MsgSpecURL":      defs.MsgSpecURL,
		"SegSpecURL":      defs.SegSpecURL,
		"DataElemSpecURL": defs.DataElemSpecURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	specSearchTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "specsearch.html")...,
	))
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

// Search term in segment specifications
func searchSegSpecs(w http.ResponseWriter, searchTerm string) []*ssp.SegSpec {
	result := ssp.SegSpecs{}
	for _, segID := range defs.Validator.SegSpecs.Ids() {
		segSpec := defs.Validator.SegSpecs.Get(segID)
		if strings.Contains(strings.ToLower(segID), searchTerm) || strings.Contains(strings.ToLower(segSpec.Name), searchTerm) {
			result = append(result, segSpec)
		}
	}
	sort.Sort(result)
	return result
}

// Search term in composite data element specifications
func searchCompositeDataElemSpecs(w http.ResponseWriter, searchTerm string) []*dsp.CompositeDataElemSpec {
	result := dsp.CompositeDataElemSpecs{}
	for _, dataElemSpec := range defs.SpecParser.CompositeDataElemSpecs {
		id := dataElemSpec.Id()
		if strings.Contains(strings.ToLower(id), searchTerm) || strings.Contains(strings.ToLower(dataElemSpec.Name()), searchTerm) {
			result = append(result, dataElemSpec)
		}
	}
	sort.Sort(result)
	return result
}

// Search term in composite data element specifications
func searchSimpleDataElemSpecs(w http.ResponseWriter, searchTerm string) []*dsp.SimpleDataElemSpec {
	result := dsp.SimpleDataElemSpecs{}
	for _, dataElemSpec := range defs.SpecParser.SimpleDataElemSpecs {
		id := dataElemSpec.Id()
		if strings.Contains(strings.ToLower(id), searchTerm) || strings.Contains(strings.ToLower(dataElemSpec.Name()), searchTerm) {
			result = append(result, dataElemSpec)
		}
	}
	sort.Sort(result)
	return result
}

func SpecSearch(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	searchTerm := strings.ToLower(r.FormValue("searchterm"))
	data := map[string]interface{}{}
	//log.Printf("Found %d message specs for search term %s", len(msgSpecs), searchTerm)
	data["msgSpecs"] = searchMsgSpecs(w, searchTerm)
	data["segSpecs"] = searchSegSpecs(w, searchTerm)
	data["compositeDataElemSpecs"] = searchCompositeDataElemSpecs(w, searchTerm)
	data["simpleDataElemSpecs"] = searchSimpleDataElemSpecs(w, searchTerm)
	data["searchTerm"] = searchTerm

	err := specSearchTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
