package entity

import "regexp"

// CEP represents a Brazilian CEP (NNNNN-NNN).
type CEP string

// IsValid validates the CEP
func (c CEP) IsValid() bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(string(c))
}
