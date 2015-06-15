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
	msgValidationTemplates = template.Must(template.ParseFiles(
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

	nestedMsg, err := defs.Validator.Validate(rawMsg)
	if err != nil {
		fmt.Fprintf(w, "Message validation failed: %s", err)
		return
	}

	fmt.Fprintf(w, "Nested msg: %s", nestedMsg.Dump())
}

func MsgValidation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := msgValidationTemplates.ExecuteTemplate(w, "layout", nil)
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
