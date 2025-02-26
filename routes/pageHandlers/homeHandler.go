package pageHandlers

import (
	"fmt"
	"net/http"

	"golang-api-lite/jsend"
	"golang-api-lite/templates"
)

// Handler
func HandleHome(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		htmlText, err := templates.GetTemplate("index.html")
		if err != nil {
			jsend.Error(w, "Error reading template", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, htmlText)
	}
}

// Handler
func HandleComponent(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		htmlText, err := templates.GetTemplate("component.html")
		if err != nil {
			jsend.Error(w, "Error reading template", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, htmlText)
	}
}
