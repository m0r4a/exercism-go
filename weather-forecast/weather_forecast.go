// Package weather forecasts the current weather
// condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current condition of the weather in the city.
var CurrentCondition string

// CurrentLocation represents the location of the city that the weather is checked.
var CurrentLocation string

// Forecast returns a string value containing the location and the weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
