package handlers

import (
	"fmt"
	"github.com/bbiskup/edify-web/defs"
	"github.com/bbiskup/edify/edifact/rawmsg"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var msgValidationTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"MsgSpecURLForId": MsgSpecURLForId,
	}
	t := template.New("layout.html").Funcs(funcMap)
	msgValidationTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "msgvalidation.html")...,
	))

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

	data := map[string]interface{}{}

	nestedMsg, err := defs.Validator.Validate(rawMsg)
	if err != nil {
		//fmt.Fprintf(w, "Message validation failed: %s", err)
		data["validationError"] = err
	} else {
		data["nestedMsg"] = nestedMsg
	}

	// fmt.Fprintf(w, "Nested msg: %s", nestedMsg.Dump())

	renderTemplate(w, data)
}

func renderTemplate(w http.ResponseWriter, data map[string]interface{}) {
	err := msgValidationTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}

func MsgValidationGET(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, nil)
}

func MsgValidationPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//log.Printf("Form: %s", r.Form)
	message := r.FormValue("message")
	validateMsg(message, w)
}
