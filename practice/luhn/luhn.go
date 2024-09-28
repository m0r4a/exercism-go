// package luhn provides a function to check if a string is valid
// according to the luhn algorithm.
package luhn

import (
	"strings"
	"unicode"
)

// Valid receives an ID as an input and ouputs weather the id
// is valid or not in the form of a true or a false.
func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	sum := calculateSum(id)

	return sum%10 == 0
}

func calculateSum(numbers string) int {
	var digitsSum int
	double := len(numbers)%2 == 0

	for _, number := range numbers {
		if !unicode.IsDigit(number) {
			return 1
		}

		digit := int(number - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		digitsSum += digit
		double = !double
	}

	return digitsSum
}
