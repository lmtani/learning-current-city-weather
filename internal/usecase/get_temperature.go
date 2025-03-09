package usecase

import (
	"fmt"
	"time"

	"github.com/lmtani/learning-current-city-weather/internal/entity"
)

// TemperatureOutputDTO represents the temperature in different units.
type TemperatureOutputDTO struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

// GetTemperature provides the temperature of a city.
type GetTemperature struct {
	weatherAPI  entity.WeatherService
	cepAPI      entity.CepService
	TimeToSleep time.Duration
}

// NewGetTemperature creates a new GetTemperature.
func NewGetTemperature(weatherAPI entity.WeatherService, cepAPI entity.CepService) *GetTemperature {
	return &GetTemperature{weatherAPI: weatherAPI, cepAPI: cepAPI, TimeToSleep: 2 * time.Second}
}

// Execute returns the temperature of the given city.
func (g *GetTemperature) Execute(cep string) (TemperatureOutputDTO, error) {
	city, err := g.retryGetCity(cep)
	if err != nil {
		return TemperatureOutputDTO{}, err
	}

	celsius, err := g.weatherAPI.GetTemperature(city)
	if err != nil {
		return TemperatureOutputDTO{}, entity.ErrWeatherAPI
	}

	t := entity.Temperature{
		Celsius: celsius,
	}

	dto := TemperatureOutputDTO{
		City:       city,
		Celsius:    t.GetCelsius(),
		Fahrenheit: t.GetFahrenheit(),
		Kelvin:     t.GetKelvin(),
	}
	return dto, nil
}

// retryGetCity retries the Get method in cepAPI up to 3 times.
func (g *GetTemperature) retryGetCity(cep string) (string, error) {
	var city string
	var err error
	for i := 0; i < 3; i++ {
		city, err = g.cepAPI.Get(cep)
		if err == nil {
			return city, nil
		}
		if err == entity.ErrCEPNotFound {
			return "", entity.ErrCEPNotFound
		}
		time.Sleep(g.TimeToSleep) // wait for 2 seconds before retrying
	}

	fmt.Println("Failed to get city after 3 retries. Assuming the city is not found.")
	return "", entity.ErrCEPNotFound
}
