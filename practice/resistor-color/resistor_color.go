// package resistorcolor provides a way to remember all the
// possible colors for a resistor and calculate its value.
package resistorcolor

var colorList = [10]string{
	"black", "brown", "red",
	"orange", "yellow", "green",
	"blue", "violet", "grey", "white",
}

// Colors returns the list of all colors.
func Colors() []string {
	return colorList[:]
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for index, value := range colorList {
		if value == color {
			return index
		}
	}

	return -1
}
