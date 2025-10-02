package pow

import (
	"fmt"
	"math"
)

// Pow computes the power of a given base raised to an exponent (base^exponent).
func Pow() (float64, error) {

	var baseInput, exponentInput float64

	fmt.Print("[ Power Operation ] Enter the base: ")
	_, err := fmt.Scan(&baseInput)
	if err != nil {
		return 0, fmt.Errorf("failed to read base input: %v", err)
	}

	fmt.Print("[ Power Operation ] Enter the exponent: ")
	_, err = fmt.Scan(&exponentInput)
	if err != nil {
		return 0, fmt.Errorf("failed to read exponent input: %v", err)
	}

	base := baseInput
	exponent := exponentInput

	// Handle special cases
	if base == 0 && exponent < 0 {
		return 0, fmt.Errorf("cannot compute 0 raised to a negative power")
	}

	return math.Pow(base, exponent), nil
}
