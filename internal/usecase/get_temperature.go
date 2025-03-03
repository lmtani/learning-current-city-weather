package usecase

import (
	"github.com/lmtani/learning-current-city-weather/internal/entity"
)

// TemperatureOutputDTO represents the temperature in different units.
type TemperatureOutputDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

// GetTemperature provides the temperature of a city.
type GetTemperature struct {
	weatherAPI entity.WeatherService
	cepAPI     entity.CepService
}

// NewGetTemperature creates a new GetTemperature.
func NewGetTemperature(weatherAPI entity.WeatherService, cepAPI entity.CepService) *GetTemperature {
	return &GetTemperature{weatherAPI: weatherAPI, cepAPI: cepAPI}
}

// Execute returns the temperature of the given city.
func (g *GetTemperature) Execute(cep string) (TemperatureOutputDTO, error) {
	c := entity.CEP(cep)
	if !c.IsValid() {
		return TemperatureOutputDTO{}, entity.ErrCEPInvalid
	}

	city, err := g.cepAPI.Get(cep)
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
		Celsius:    t.GetCelsius(),
		Fahrenheit: t.GetFahrenheit(),
		Kelvin:     t.GetKelvin(),
	}
	return dto, nil
}
