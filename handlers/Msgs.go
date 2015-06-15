package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	msp "github.com/bbiskup/edify/edifact/spec/message"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var msgsTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"MsgSpecURL": MsgSpecURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	msgsTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths(
			"layout.html", "navbar.html",
			"msgspecs.html", "msgspectable.html")...,
	))
}

func MsgSpecs(w http.ResponseWriter, r *http.Request) {
	msgSpecs := make(msp.MsgSpecs, 0, len(defs.Validator.MsgSpecs))
	for _, msgSpec := range defs.Validator.MsgSpecs {
		msgSpecs = append(msgSpecs, msgSpec)
	}
	sort.Sort(msgSpecs)

	err := msgsTemplates.ExecuteTemplate(w, "layout", msgSpecs)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
