package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var msgSpecTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"MsgSpecURL":     MsgSpecURL,
		"MsgSpecPartURL": MsgSpecPartURL,
		"MsgSpecGrpURL":  MsgSpecGrpURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	msgSpecTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths(
			"layout.html", "navbar.html", "msgspec.html", "msgspecgrpchildren.html")...,
	))
}

func MsgSpec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msgSpecID := vars["id"]
	msgSpec := defs.Validator.MsgSpecs[msgSpecID]

	data := map[string]interface{}{
		"msgSpec":  msgSpec,
		"children": msgSpec.TopLevelGrp.Children(),
	}

	err := msgSpecTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
