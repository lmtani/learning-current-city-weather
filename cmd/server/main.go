package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/lmtani/learning-current-city-weather/internal/entity"
	"github.com/lmtani/learning-current-city-weather/internal/infra/cep"
	"github.com/lmtani/learning-current-city-weather/internal/usecase"
	"github.com/lmtani/learning-current-city-weather/pkg/weather"
)

func main() {
	fmt.Println("Server is running on port 8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetTemperature)

	// Wrap the mux with the logging middleware.
	loggedMux := handlers.LoggingHandler(os.Stdout, mux)

	if err := http.ListenAndServe("0.0.0.0:8080", loggedMux); err != nil {
		log.Fatal(err)
	}
}

// GetTemperature returns the temperature of a city.
func GetTemperature(w http.ResponseWriter, r *http.Request) {
	cepService := cep.NewService()
	weatherService := weather.NewService(os.Getenv("WEATHER_API_KEY"))
	getTemperature := usecase.NewGetTemperature(weatherService, cepService)
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "Missing 'cep' query parameter", http.StatusBadRequest)
		return
	}

	output, err := getTemperature.Execute(cep)
	if err != nil {
		switch err {
		case entity.ErrCEPInvalid:
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		case entity.ErrCEPNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Return the temperature in JSON format.
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
