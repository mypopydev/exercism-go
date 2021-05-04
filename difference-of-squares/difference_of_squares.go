package diffsquares

// SumOfSquares get the sum of the squares
// SumOfSquares =  n * (n+1) * (2n+1)/6
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSum get the the square of the sum
// SquareOfSum = ((n * (n+1))/2)^2
func SquareOfSum(n int) int {
	return ((n * (n + 1)) * (n * (n + 1))) / 4
}

// Difference find the difference between the square of the sum and
// the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
