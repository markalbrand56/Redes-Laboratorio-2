package main

import (
	"testing"
)

func TestHammingDecode(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"0110011", "0110"},
		{"1100001", "1100"},
		{"110011111", "10011"},
		{"0111001010", "011000"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			decoded := HammingDecode(test.input)
			if decoded != test.output {
				t.Errorf("For message %s, expected %s but got %s", test.input, test.output, decoded)
			}
		})
	}
}
