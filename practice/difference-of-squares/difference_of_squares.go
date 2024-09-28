// package diffsquares is a package to check the difference
// between the sum of the first ten natural numbers and square them
// and the first ten squared numbers and sum them.
package diffsquares

// SquareOfSum is a function that recieves a number n and returns the
// value of the n numbers summed and then squared.
func SquareOfSum(n int) int {
	var sum int

	for i := 1; i < (n + 1); i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares is a function that recieves a number n and returns the
// value of the n numbers summed and the power of 2 and summing each.
func SumOfSquares(n int) int {
	var sum int

	for i := 1; i < (n + 1); i++ {
		sum += i * i
	}

	return sum
}

// Difference is a function that recieves a number n and returns the
// difference between summing all the elements and then square them
// and square each element and sum eachother
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
