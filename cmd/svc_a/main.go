package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/lmtani/learning-current-city-weather/internal/entity"
	"github.com/lmtani/learning-current-city-weather/internal/usecase"
)

func main() {
	fmt.Println("Service A is running on port 8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleCEPInput)
	// Wrap the mux with the logging middleware.
	loggedMux := handlers.LoggingHandler(os.Stdout, mux)

	if err := http.ListenAndServe("0.0.0.0:8080", loggedMux); err != nil {
		log.Fatal(err)
	}
}

func handleCEPInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		CEP string `json:"cep"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	c := entity.CEP(payload.CEP)
	if !c.IsValid() {
		http.Error(w, entity.ErrCEPInvalid.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Call Service B
	resp, err := http.Get("http://service-b:8080/?cep=" + payload.CEP)
	if err != nil {
		http.Error(w, "Failed to get city", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var city usecase.TemperatureOutputDTO
	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(&city); err != nil {
			http.Error(w, "Failed to decode response from service-b", http.StatusInternalServerError)
			return
		}
	case http.StatusNotFound:
		http.Error(w, entity.ErrCEPNotFound.Error(), http.StatusNotFound)
		return

	default:
		http.Error(w, "Failed to get city", http.StatusInternalServerError)
		return
	}

	// Return the city in JSON format.
	err = json.NewEncoder(w).Encode(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
