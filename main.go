package main

import (
	"github.com/bbiskup/edify-web/defs"
	"github.com/bbiskup/edify-web/handlers"

	//"github.com/gorilla/mux"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	static := http.FileServer(http.Dir(defs.STATIC_DIR))
	bower := http.FileServer(http.Dir(defs.BOWER_DIR))

	mux.Handle("/static/edify/", http.StripPrefix("/static/edify/", static))
	mux.Handle("/static/bower/", http.StripPrefix("/static/bower/", bower))
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/about/", handlers.About)

	server := &http.Server{
		Addr:    "0.0.0.0:8001",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
