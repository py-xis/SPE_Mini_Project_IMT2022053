package fact

import (
	"fmt"
	"os"
	"testing"
)

func TestFact(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		expectErr bool
		expected  int64
	}{
		{
			name:      "Zero factorial",
			input:     0,
			expectErr: false,
			expected:  1,
		},
		{
			name:      "One factorial",
			input:     1,
			expectErr: false,
			expected:  1,
		},
		{
			name:      "Small positive number",
			input:     5,
			expectErr: false,
			expected:  120,
		},
		{
			name:      "Larger positive number",
			input:     10,
			expectErr: false,
			expected:  3628800,
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

			result, err := Fact()
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

// mockInput simulates user input for testing purposes.
func mockInput(input int) *os.File {
	r, w, _ := os.Pipe()
	fmt.Fprintln(w, input)
	w.Close()
	return r
}
