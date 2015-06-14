package main

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/bbiskup/edify-web/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	static := http.FileServer(http.Dir(defs.STATIC_DIR))
	bower := http.FileServer(http.Dir(defs.BOWER_DIR))

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/about/", handlers.About)

	r.PathPrefix("/static/edify/").Handler(http.StripPrefix("/static/edify/", static))
	r.PathPrefix("/static/bower/").Handler(http.StripPrefix("/static/bower/", bower))

	server := &http.Server{
		Addr:    "0.0.0.0:8001",
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
