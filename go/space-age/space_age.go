package space

import "strings"

type Planet string

func Age(seconds float64, planet Planet) float64 {
	secondsInYear := 1.0
	switch strings.ToLower(string(planet)) {
	case "earth":
		secondsInYear = 31557600.0
	case "mercury":
		secondsInYear = 31557600.0 * 0.2408467
	case "venus":
		secondsInYear = 31557600.0 * 0.61519726
	case "mars":
		secondsInYear = 31557600.0 * 1.8808158
	case "jupiter":
		secondsInYear = 31557600.0 * 11.862615
	case "saturn":
		secondsInYear = 31557600.0 * 29.447498
	case "uranus":
		secondsInYear = 31557600.0 * 84.016846
	case "neptune":
		secondsInYear = 31557600.0 * 164.79132
	default:
		seconds = -1
	}
	return seconds / secondsInYear
}
