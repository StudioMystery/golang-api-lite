package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-api-lite/data"
	"golang-api-lite/jsend"
	"golang-api-lite/library"
)

// Handler
func HandleFooBar(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if !req.URL.Query().Has("Id") {
			jsend.Error(w, "Must have Id", http.StatusBadRequest)
			return
		}
		getFooBar(w, req)
	case http.MethodPost:
		postFooBar(w, req)
	case http.MethodPut:
		putFooBar(w, req)
	case http.MethodDelete:
		if !req.URL.Query().Has("Id") {
			jsend.Error(w, "Must have Id", http.StatusBadRequest)
			return
		}
		deleteFooBar(w, req)
	default:
		jsend.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Internal Functions

func getFooBar(w http.ResponseWriter, req *http.Request) {
	fooBarId, err := strconv.Atoi(req.URL.Query().Get("Id"))
	if err != nil {
		jsend.Error(w, "Error reading Id", http.StatusBadRequest)
		return
	}

	result, err := data.ReadFooBar(fooBarId)

	if err != nil {
		jsend.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsend.Success(w, result, http.StatusOK)
}

func postFooBar(w http.ResponseWriter, req *http.Request) {
	var dto library.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		jsend.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	result, err := data.CreateFooBar(dto)
	if err != nil {
		jsend.Error(w, "Error saving data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsend.Success(w, result, http.StatusOK)
}

func putFooBar(w http.ResponseWriter, req *http.Request) {
	var dto library.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		jsend.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	rows, err := data.UpdateFooBar(dto)
	if err != nil {
		jsend.Error(w, "Error saving data", http.StatusInternalServerError)
		return
	}

	if rows == 0 {
		jsend.Error(w, "No rows affected", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsend.Success(w, rows, http.StatusOK)
}

func deleteFooBar(w http.ResponseWriter, req *http.Request) {
	fooBarId, err := strconv.Atoi(req.URL.Query().Get("Id"))
	if err != nil {
		jsend.Error(w, "Error reading Id", http.StatusBadRequest)
		return
	}

	err = data.DeleteFooBar(fooBarId)
	if err != nil {
		jsend.Error(w, "Error deleting data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsend.Success(w, nil, http.StatusOK)
}
