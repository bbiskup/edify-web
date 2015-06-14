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

	SPEC_DIR = ".edify/downloads/d14b"
)

var indexTemplates *template.Template
var aboutTemplates *template.Template
var validator *validation.MsgValidator

func init() {
	// indexTemplates = template.Must(template.ParseFiles(
	// 	templatePaths("layout.html", "navbar.html", "content.html")...,
	// ))

	indexTemplates = template.Must(template.ParseFiles(
		templatePaths("layout.html", "navbar.html", "content.html")...,
	))

	aboutTemplates = template.Must(template.ParseFiles(
		templatePaths("layout.html", "navbar.html", "about.html")...,
	))

	var err error
	validator, err = validation.GetMsgValidator("14B", SPEC_DIR)
	if err != nil {
		panic(fmt.Sprintf("Unable to create validator: %s", err))
	}
}

func templatePaths(templateNames ...string) []string {
	result := make([]string, 0, len(templateNames))
	for _, templateName := range templateNames {
		result = append(result, TPL_DIR+string(os.PathSeparator)+templateName)
	}
	return result
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

func index(w http.ResponseWriter, r *http.Request) {
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

func about(w http.ResponseWriter, r *http.Request) {
	err := aboutTemplates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}

func main() {
	mux := http.NewServeMux()
	static := http.FileServer(http.Dir(STATIC_DIR))
	bower := http.FileServer(http.Dir(BOWER_DIR))

	mux.Handle("/static/edify/", http.StripPrefix("/static/edify/", static))
	mux.Handle("/static/bower/", http.StripPrefix("/static/bower/", bower))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/about/", about)

	server := &http.Server{
		Addr:    "0.0.0.0:8001",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
