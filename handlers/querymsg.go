package handlers

import (
	"fmt"
	"github.com/bbiskup/edify-web/defs"
	"github.com/bbiskup/edify/edifact/query"
	"github.com/bbiskup/edify/edifact/rawmsg"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var queryMsgTemplates *template.Template

func init() {
	funcMap := template.FuncMap{
		"MsgSpecURLForId": MsgSpecURLForId,
	}
	t := template.New("layout.html").Funcs(funcMap)
	queryMsgTemplates = template.Must(t.ParseFiles(
		defs.TemplatePaths("layout.html", "navbar.html", "querymsg.html")...,
	))

}

func queryMsg(message string, queryStr string, w http.ResponseWriter) {
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
		navigator := query.NewNavigator()
		queryResult, err := navigator.Navigate(queryStr, nestedMsg)
		if err != nil {
			data["validationError"] = err
		}
		data["queryResult"] = queryResult
	}

	// fmt.Fprintf(w, "Nested msg: %s", nestedMsg.Dump())

	renderQueryTemplate(w, data)
}

func renderQueryTemplate(w http.ResponseWriter, data map[string]interface{}) {
	err := queryMsgTemplates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
	}
}

func QueryMsgGET(w http.ResponseWriter, r *http.Request) {
	renderQueryTemplate(w, nil)
}

func QueryMsgPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//log.Printf("Form: %s", r.Form)
	message := r.FormValue("message")
	queryStr := r.FormValue("query")
	queryMsg(message, queryStr, w)
}
