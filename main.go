package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	STATIC_DIR = "static"
	BOWER_DIR  = "bower_components"
	//  TPL_DIR    = strings.Join([]string{STATIC_DIR, "templates"}, string(os.PathSeparator))

	TPL_DIR = STATIC_DIR + string(os.PathSeparator) + "templates"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob(
		TPL_DIR + string(os.PathSeparator) + "*",
	))
}

func index(w http.ResponseWriter, request *http.Request) {
	err := templates.ExecuteTemplate(w, "layout", nil)
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

	server := &http.Server{
		Addr:    "0.0.0.0:8001",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
