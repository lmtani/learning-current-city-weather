package entity

import "context"

// WeatherService provides the temperature of a city.
type WeatherService interface {
	GetTemperature(ctx context.Context, city string) (float64, error)
}

// CepService provides the city of a given CEP.
type CepService interface {
	Get(ctx context.Context, queryCEP string) (string, error)
}
