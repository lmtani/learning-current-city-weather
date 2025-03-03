package entity

// Temperature represents the temperature of a city.
type Temperature struct {
	Celsius float64
}

// GetKelvin returns the temperature in Kelvin.
func (w Temperature) GetKelvin() float64 {
	return w.Celsius + 273
}

// GetCelsius returns the temperature in Celsius.
func (w Temperature) GetCelsius() float64 {
	return w.Celsius
}

// GetFahrenheit returns the temperature in Fahrenheit.
func (w Temperature) GetFahrenheit() float64 {
	return w.Celsius*1.8 + 32
}
