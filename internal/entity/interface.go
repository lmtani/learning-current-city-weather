package entity

// WeatherService provides the temperature of a city.
type WeatherService interface {
	GetTemperature(city string) (float64, error)
}

// CepService provides the city of a given CEP.
type CepService interface {
	Get(queryCEP string) (string, error)
}
