package calc

import (
	"errors"

	"github.com/fatih/color"
)

// Add returns the sum of passed in arguments with error
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		errMsg := color.RedString("provide more than 2 numbers")
		return sum, errors.New(errMsg)
	}

	for _, number := range numbers {
		sum += number
	}
	return sum, nil
}
