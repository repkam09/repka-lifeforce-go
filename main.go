package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/repkam09/repka-lifeforce-go/handlers"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc(handlers.HelloHandlerURL, handlers.HelloHandler).Methods("GET")
	router.HandleFunc("/api/election/full", handlers.ElectionResultHandler).Methods("GET")
	router.HandleFunc("/api/weather/current/zip/{zip}", handlers.WeatherZipCodeCurrentHandler).Methods("GET")
	router.HandleFunc("/api/weather/forecast/zip/{zip}", handlers.WeatherZipCodeForecastHandler).Methods("GET")

	log.Println("Server is running on port 16001")
	log.Fatal(http.ListenAndServe(":16001", router))
}
