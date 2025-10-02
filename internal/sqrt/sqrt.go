package sqrt

import (
	"fmt"
	"math"
)

// Sqrt computes the square root of a given number using the math.Sqrt function.
func Sqrt() (float64, error) {

	var xInput float64

	fmt.Print("[ Sqrt Operation ] Enter a number: ")

	_, err := fmt.Scan(&xInput)
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %v", err)
	}

	x := xInput
	if x < 0 {
		return 0, fmt.Errorf("cannot compute square root of a negative number: %f", x)
	}

	return math.Sqrt(x), nil
}