// package leap provides a function to let you know if
// the you write is leap or not.
package leap

// IsLeapYear recieves a year in a form of int and returns you
// a boolean value if it's leap or not.
func IsLeapYear(year int) bool {
	if (year%4) == 0 && (year%100) != 0 {
		return true

	} else if (year%100) == 0 && (year%400) == 0 {
		return true

	}

	return false
}
