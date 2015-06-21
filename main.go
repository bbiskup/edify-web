package main

import (
	"fmt"
	"github.com/bbiskup/edify-web/defs"
	"github.com/bbiskup/edify-web/handlers"
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	hostFlag = cli.StringFlag{
		Name:  "host, H",
		Value: "localhost",
		Usage: "Host name or IP address of web server",
	}
	portFlag = cli.IntFlag{
		Name:  "port, p",
		Value: 8001,
		Usage: "TCP port of web server",
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "edify-web"
	app.Usage = "EDIFACT web application"
	app.EnableBashCompletion = true

	r := mux.NewRouter()
	static := http.FileServer(http.Dir(defs.StaticDir))
	bower := http.FileServer(http.Dir(defs.BowerDir))

	r.HandleFunc("/", handlers.Index).
		Name("index").
		Methods("GET")

	r.HandleFunc("/message/query/", handlers.QueryMsgGET).
		Name("querysg").
		Methods("GET")
	r.HandleFunc("/message/query/", handlers.QueryMsgPOST).
		Name("querysg").
		Methods("POST")

	r.HandleFunc("/message/validation/", handlers.MsgValidationGET).
		Name("messagevalidation").
		Methods("GET")
	r.HandleFunc("/message/validation/", handlers.MsgValidationPOST).
		Name("messagevalidation").
		Methods("POST")

	r.HandleFunc("/specsearch/", handlers.SpecSearch).
		Name("specsearch").
		Methods("GET")

	r.HandleFunc("/browsespecs/", handlers.BrowseSpecs).
		Name("browsespecs").
		Methods("GET")

	r.HandleFunc("/about/", handlers.About).
		Name("about").
		Methods("GET")

	r.HandleFunc("/specs/message/{id}/{segGroupName}", handlers.MsgSpecSegGrp).
		Name("msgspecseggrp").
		Methods("GET")
	r.HandleFunc("/specs/message/{id}", handlers.MsgSpec).
		Name("msgspec").
		Methods("GET")
	r.HandleFunc("/specs/message/", handlers.MsgSpecs).
		Name("message").
		Methods("GET")

	r.HandleFunc("/specs/segment/{id}", handlers.SegSpec).
		Name("segspecs").
		Methods("GET")
	r.HandleFunc("/specs/segment/", handlers.SegSpecs).
		Name("segments").
		Methods("GET")

	r.HandleFunc("/specs/compositedataelement/{id}", handlers.CompositeDataElemSpec).
		Name("compositedataelemspec").
		Methods("GET")
	r.HandleFunc("/specs/compositedataelement/", handlers.CompositeDataElemSpecs).
		Name("compositedataelements").
		Methods("GET")

	r.HandleFunc("/specs/simpledataelement/{id}", handlers.SimpleDataElemSpec).
		Name("simpledataelemspec").
		Methods("GET")
	r.HandleFunc("/specs/simpledataelement/", handlers.SimpleDataElemSpecs).
		Name("simpledataelements").
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

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Usage:   "Run web application",
			Aliases: []string{"r"},
			Action: func(c *cli.Context) {
				addrStr := fmt.Sprintf("%s:%d", c.String("host"), c.Int("port"))
				log.Printf("Listening on %s", addrStr)
				server := &http.Server{
					Addr:    addrStr,
					Handler: r,
				}
				if err := server.ListenAndServe(); err != nil {
					panic(err)
				}
			},
			Flags: []cli.Flag{
				hostFlag, portFlag,
			},
		},
	}

	app.Run(os.Args)
}
