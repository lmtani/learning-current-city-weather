package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Weather represents the weather data.
type Weather struct {
	Celsius    float64 `json:"temp_c"`
	Fahrenheit float64 `json:"temp_f"`
}

// API is a weather service.
type API struct {
	apiKey string
}

// NewService creates a new weather service.
func NewService(apiKey string) *API {
	return &API{apiKey: apiKey}
}

// GetTemperature returns the temperature of the given city.
func (a *API) GetTemperature(city string) (float64, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s&key=%s", city, a.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	var result struct {
		Current Weather `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0.0, err
	}

	return result.Current.Celsius, nil
}
