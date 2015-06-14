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
	funcMap := template.FuncMap{"msgSpecURL": msgSpecURL}
	t := template.New("layout.html").Funcs(funcMap)
	msgSpecTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "msgspec.html")...,
	))
}

func MsgSpec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msgSpecID := vars["id"]

	data := map[string]interface{}{
		"msgSpec": defs.Validator.MsgSpecs[msgSpecID],
	}

	err := msgSpecTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
