package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/repkam09/repka-lifeforce-go/utils/common"
	"github.com/repkam09/repka-lifeforce-go/utils/config"
)

type WeatherZipCodeCurrentResponse struct {
	Message string `json:"message"`
}

type WeatherZipCodeForecastResponse struct {
	Message string `json:"message"`
}

func WeatherZipCodeCurrentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	zipCode := vars["zip"]
	url := WeatherZipCodeCurrentBuilder(zipCode)
	common.ForwardResponse(url, w)
}

func WeatherZipCodeForecastHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	zipCode := vars["zip"]
	url := WeatherZipCodeForecastBuilder(zipCode)
	common.ForwardResponse(url, w)
}

func WeatherZipCodeCurrentBuilder(zipCode string) string {
	return "http://api.openweathermap.org/data/2.5/weather?zip=" + zipCode + "&appid=" + config.GetConfig().WEATHER_API_KEY
}

func WeatherZipCodeForecastBuilder(zipCode string) string {
	return "http://api.openweathermap.org/data/2.5/forecast?zip=" + zipCode + "&appid=" + config.GetConfig().WEATHER_API_KEY
}
