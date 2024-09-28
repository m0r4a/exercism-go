// package that allows you to calculate your age in
// certain planets upon seconds.
package space

type Planet string

const secondsInYear = 31557600

// Planets map with the values of the equivalents of
// Earth year ages.
var planets = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// function Age recieves your age in seconds and the desired planet
// and returns a float64 containing your age in years in the
// planet you chose.
func Age(seconds float64, planet Planet) float64 {
	if !inPlanets(planet) {
		return -1
	}

	return seconds / secondsInYear / planets[planet]
}

// function inPlanets is a function that
// checks if certain planet exists in the
// planets dictionary.
func inPlanets(planet Planet) bool {
	_, ok := planets[planet]
	if !ok {
		return false
	}

	return true
}
