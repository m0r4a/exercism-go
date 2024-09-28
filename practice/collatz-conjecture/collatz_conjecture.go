// package collatzconjecture gives you the amount of iterations
// needed to reach one in the Collatz Conjecture (3x + 1 problem)
package collatzconjecture

import "errors"

// The CollatzConjecture function reciveces an n number and returns
// the amount of iterations needed to reach a value of 1
func CollatzConjecture(n int) (int, error) {
	var iterations int = 0

	if n <= 0 {
		return iterations, errors.New("value of n must be a positive integer")
	}

	for n != 1 {

		if n%2 == 0 {
			n /= 2
		} else {
			n = (3 * n) + 1
		}

		iterations += 1
	}

	return iterations, nil
}
