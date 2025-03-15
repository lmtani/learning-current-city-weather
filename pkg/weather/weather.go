package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

// Weather represents the weather data.
type Weather struct {
	Celsius    float64 `json:"temp_c"`
	Fahrenheit float64 `json:"temp_f"`
}

// API is a weather service.
type API struct {
	apiKey string
	Tracer trace.Tracer
}

// NewService creates a new weather service.
func NewService(apiKey string) *API {
	return &API{apiKey: apiKey}
}

// GetTemperature returns the temperature of the given city.
func (a *API) GetTemperature(ctx context.Context, city string) (float64, error) {
	// if tracer is provided, use it
	if a.Tracer != nil {
		_, span := a.Tracer.Start(ctx, "request-weather-api")
		defer span.End()
	}

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
