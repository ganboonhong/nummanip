package calc

// Add returns the sum of passed in arguments
func Add(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
