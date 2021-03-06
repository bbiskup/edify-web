package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	csp "github.com/bbiskup/edify/edifact/spec/codes"
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
		"MsgSpecURL":      MsgSpecURL,
		"SegSpecURL":      SegSpecURL,
		"DataElemSpecURL": DataElemSpecURL,
		"CodesSpecURL":    CodesSpecURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	specSearchTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths(
			"layout.html", "navbar.html",
			"specsearch.html", "msgspectable.html")...,
	))
}

// Search term in message specifications
func searchMsgSpecs(searchTerm string) []*msp.MsgSpec {
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
func searchSegSpecs(searchTerm string) []*ssp.SegSpec {
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
func searchCompositeDataElemSpecs(searchTerm string) []*dsp.CompositeDataElemSpec {
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

// Search term in simple data element specifications
func searchSimpleDataElemSpecs(searchTerm string) []*dsp.SimpleDataElemSpec {
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

// Search term in code specs
func searchCodesSpecs(searchTerm string) []*csp.CodesSpec {
	result := []*csp.CodesSpec{}
	for _, codesSpec := range defs.SpecParser.CodesSpecs {
		id := codesSpec.Id
		if strings.Contains(strings.ToLower(id), searchTerm) || strings.Contains(strings.ToLower(codesSpec.Name), searchTerm) {
			result = append(result, codesSpec)
		}
	}
	return result
}

func SpecSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchTerm := strings.ToLower(r.FormValue("searchterm"))
	data := map[string]interface{}{}
	//log.Printf("Found %d message specs for search term %s", len(msgSpecs), searchTerm)
	data["msgSpecs"] = searchMsgSpecs(searchTerm)
	data["segSpecs"] = searchSegSpecs(searchTerm)
	data["compositeDataElemSpecs"] = searchCompositeDataElemSpecs(searchTerm)
	data["simpleDataElemSpecs"] = searchSimpleDataElemSpecs(searchTerm)
	data["codesSpecs"] = searchCodesSpecs(searchTerm)
	data["searchTerm"] = searchTerm

	err := specSearchTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
