// Package weather handles displaying the weather forecast to the user.
package weather

// CurrentCondition represents a string to display condition in a city.
var CurrentCondition string

// CurrentLocation represents a string to display the location via city.
var CurrentLocation string

// Forecast function to take city, and condition and display it to the user.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
