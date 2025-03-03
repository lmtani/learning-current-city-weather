package entity

import "errors"

var (
	ErrCEPInvalid  = errors.New("invalid zipcode")
	ErrCEPNotFound = errors.New("can not find zipcode")
	ErrWeatherAPI  = errors.New("weather service error")
)
