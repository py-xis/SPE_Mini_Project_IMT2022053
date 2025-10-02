package ln

import (
	"fmt"
	"math"
)

// Ln computes the natural logarithm of a given positive number.
func Ln() (float64, error) {

	var xInput float64

	fmt.Print("[ Natural Logarithm Operation ] Enter a positive number: ")

	_, err := fmt.Scan(&xInput)
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %v", err)
	}

	x := xInput
	if x <= 0 {
		return 0, fmt.Errorf("cannot compute natural logarithm of a non-positive number: %f", x)
	}

	return math.Log(x), nil
}
