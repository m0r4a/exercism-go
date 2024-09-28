// package hamming provides a function in that calculates
// the hamming distance between two DNA strands.
package hamming

import "errors"

// Distance recieves two strings and returns the hamming
// distance between them
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b must be the same lenght")
	}

	var distance int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance, nil
}
