package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

var HelloHandlerURL string = "/"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "This is a simple API that provides weather and election information."}
	json.NewEncoder(w).Encode(response)
}
