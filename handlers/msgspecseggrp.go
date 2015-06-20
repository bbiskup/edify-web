package handlers

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var msgSpecSegGrpTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"MsgSpecURL":     MsgSpecURL,
		"MsgSpecPartURL": MsgSpecPartURL,
		"MsgSpecGrpURL":  MsgSpecGrpURL,
	}
	t := template.New("layout.html").Funcs(funcMap)
	msgSpecSegGrpTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths(
			"layout.html", "navbar.html", "msgspecseggrp.html", "msgspecgrpchildren.html")...,
	))
}

func MsgSpecSegGrp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msgSpecID := vars["id"]
	msgSpec := defs.Validator.MsgSpecs[msgSpecID]
	segGrpName := vars["segGroupName"]

	segGrp, err := msgSpec.FindSegGrpSpec(segGrpName)
	if err != nil {
		log.Printf("Segment group '%s' could not be found: %s", segGrpName, err)
	}

	data := map[string]interface{}{
		"msgSpec":    defs.Validator.MsgSpecs[msgSpecID],
		"children":   segGrp.Children(),
		"segGrpName": segGrpName,
	}

	err = msgSpecSegGrpTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}
