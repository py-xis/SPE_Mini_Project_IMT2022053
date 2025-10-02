package pow

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestPow(t *testing.T) {
	tests := []struct {
		name      string
		base      float64
		exponent  float64
		expectErr bool
		expected  float64
	}{
		{
			name:      "Positive base and exponent",
			base:      2,
			exponent:  3,
			expectErr: false,
			expected:  8,
		},
		{
			name:      "Base to power of zero",
			base:      5,
			exponent:  0,
			expectErr: false,
			expected:  1,
		},
		{
			name:      "Base to power of one",
			base:      7,
			exponent:  1,
			expectErr: false,
			expected:  7,
		},
		{
			name:      "Negative exponent",
			base:      2,
			exponent:  -2,
			expectErr: false,
			expected:  0.25,
		},
		{
			name:      "Zero to negative power",
			base:      0,
			exponent:  -1,
			expectErr: true,
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock user input
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			os.Stdin = mockInput(tt.base, tt.exponent)

			result, err := Pow()
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if math.Abs(result-tt.expected) > 1e-10 {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

// mockInput simulates user input for testing purposes.
func mockInput(base, exponent float64) *os.File {
	r, w, _ := os.Pipe()
	fmt.Fprintln(w, base)
	fmt.Fprintln(w, exponent)
	w.Close()
	return r
}
