package ln

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestLn(t *testing.T) {
	tests := []struct {
		name      string
		input     float64
		expectErr bool
		expected  float64
	}{
		{
			name:      "Natural logarithm of e",
			input:     math.E,
			expectErr: false,
			expected:  1,
		},
		{
			name:      "Natural logarithm of 1",
			input:     1,
			expectErr: false,
			expected:  0,
		},
		{
			name:      "Positive number",
			input:     10,
			expectErr: false,
			expected:  math.Log(10),
		},
		{
			name:      "Zero",
			input:     0,
			expectErr: true,
			expected:  0,
		},
		{
			name:      "Negative number",
			input:     -5,
			expectErr: true,
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock user input
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			os.Stdin = mockInput(tt.input)

			result, err := Ln()
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
func mockInput(input float64) *os.File {
	r, w, _ := os.Pipe()
	fmt.Fprintln(w, input)
	w.Close()
	return r
}
