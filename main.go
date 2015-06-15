package main

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/bbiskup/edify-web/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	static := http.FileServer(http.Dir(defs.StaticDir))
	bower := http.FileServer(http.Dir(defs.BowerDir))

	r.HandleFunc("/", handlers.Index).
		Name("index").
		Methods("GET")

	r.HandleFunc("/message/validation/", handlers.MsgValidationGET).
		Name("messagevalidation").
		Methods("GET")
	r.HandleFunc("/message/validation/", handlers.MsgValidationPOST).
		Name("messagevalidation").
		Methods("POST")

	r.HandleFunc("/specsearch/", handlers.SpecSearch).
		Name("specsearch").
		Methods("GET")

	r.HandleFunc("/about/", handlers.About).
		Name("about").
		Methods("GET")

	r.HandleFunc("/specs/message/{id}", handlers.MsgSpec).
		Name("msgspec").
		Methods("GET")
	r.HandleFunc("/specs/message/", handlers.MsgSpecs).
		Name("message").
		Methods("GET")

	r.HandleFunc("/specs/segment/{id}", handlers.SegSpec).
		Name("segspec").
		Methods("GET")

	r.HandleFunc("/specs/compositedataelement/{id}", handlers.CompositeDataElemSpec).
		Name("compositedataelemspec").
		Methods("GET")

	r.HandleFunc("/specs/simpledataelement/{id}", handlers.SimpleDataElemSpec).
		Name("simpledataelemspec").
		Methods("GET")

	r.HandleFunc("/specs/code/{id}", handlers.CodeSpec).
		Name("codespec").
		Methods("GET")

	r.PathPrefix("/static/edify/").
		Name("edifystatic").
		Handler(http.StripPrefix("/static/edify/", static))

	r.PathPrefix("/static/bower/").
		Name("bowerstatic").
		Handler(http.StripPrefix("/static/bower/", bower))

	// rev, err := r.GetRoute("about").URL()
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("#### reverse: %s", rev)

	server := &http.Server{
		Addr:    "0.0.0.0:8001",
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
