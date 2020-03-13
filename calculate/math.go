package calc

import (
	"errors"
)

// Add returns the sum of passed in arguments with error
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("provide more than 2 numbers")
	}

	for _, number := range numbers {
		sum += number
	}
	return sum, nil
}
