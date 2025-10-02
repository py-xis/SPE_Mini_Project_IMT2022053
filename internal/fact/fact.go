package fact

import (
	"fmt"
)

// Fact computes the factorial of a given non-negative integer.
func Fact() (int64, error) {

	var nInput int

	fmt.Print("[ Factorial Operation ] Enter a non-negative integer: ")

	_, err := fmt.Scan(&nInput)
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %v", err)
	}

	n := nInput
	if n < 0 {
		return 0, fmt.Errorf("cannot compute factorial of a negative number: %d", n)
	}

	// Handle edge cases
	if n == 0 || n == 1 {
		return 1, nil
	}

	// Calculate factorial iteratively
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}

	return result, nil
}
