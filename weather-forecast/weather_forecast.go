// Package weather get current location, condition.
package weather

// CurrentCondition store current condition.
var CurrentCondition string

// CurrentLocation store current location.
var CurrentLocation string

// Forecast get the location, condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
