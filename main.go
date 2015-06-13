package main

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/rawmsg"
	"github.com/bbiskup/edify/edifact/validation"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	STATIC_DIR = "static"
	BOWER_DIR  = "bower_components"
	//  TPL_DIR    = strings.Join([]string{STATIC_DIR, "templates"}, string(os.PathSeparator))

	TPL_DIR = STATIC_DIR + string(os.PathSeparator) + "templates"
)

var templates *template.Template
var validator *validation.MsgValidator

func init() {
	templates = template.Must(template.ParseGlob(
		TPL_DIR + string(os.PathSeparator) + "*",
	))

	var err error
	validator, err = validation.GetMsgValidator("14B", "testdata/d14b")
	if err != nil {
		panic(fmt.Sprintf("Unable to create validator: %s", err))
	}
}

func validateMsg(message string, w http.ResponseWriter) {
	//log.Printf("Message '%s'", message)
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

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("r: %s", r.Method)
	if r.Method == "GET" {
		err := templates.ExecuteTemplate(w, "layout", nil)
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

func main() {
	mux := http.NewServeMux()
	static := http.FileServer(http.Dir(STATIC_DIR))
	bower := http.FileServer(http.Dir(BOWER_DIR))

	mux.Handle("/static/edify/", http.StripPrefix("/static/edify/", static))
	mux.Handle("/static/bower/", http.StripPrefix("/static/bower/", bower))
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8001",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
