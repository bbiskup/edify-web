package handlers

import (
	"fmt"
	"github.com/bbiskup/edify-web/defs"
	"github.com/bbiskup/edify/edifact/rawmsg"
	"github.com/bbiskup/edify/edifact/validation"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var indexTemplates *template.Template
var validator *validation.MsgValidator

func init() {
	indexTemplates = template.Must(template.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "content.html")...,
	))

	var err error
	validator, err = validation.GetMsgValidator("14B", defs.SPEC_DIR)
	if err != nil {
		panic(fmt.Sprintf("Unable to create validator: %s", err))
	}
}

func validateMsg(message string, w http.ResponseWriter) {
	log.Printf("Message '%s'", message)
	splitMsg := strings.Split(message, "\r\n")
	joinedMsg := strings.Join(splitMsg, "")
	//log.Printf("Joined msg: %#v", joinedMsg)
	var rawMsg *rawmsg.RawMsg
	rawMsgParser := rawmsg.NewParser()
	rawMsg, err := rawMsgParser.ParseRawMsg(joinedMsg)
	if err != nil {
		fmt.Fprintf(w, "Parsing raw message failed: %s", err)
		return
	}

	nestedMsg, err := validator.Validate(rawMsg)
	if err != nil {
		fmt.Fprintf(w, "Message validation failed: %s", err)
		return
	}

	fmt.Fprintf(w, "Nested msg: %s", nestedMsg.Dump())
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := indexTemplates.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			log.Printf("Error executing template: %s", err)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		//log.Printf("Form: %s", r.Form)
		message := r.FormValue("message")
		validateMsg(message, w)

	} else {
		panic(fmt.Sprintf("Unsupported method %s", r.Method))
	}
}
