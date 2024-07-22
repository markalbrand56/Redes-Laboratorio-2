package main

import (
	"testing"
)

func TestHammingDecode(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		// Cadenas sin error
		{"0110011", "0110"},
		{"1100001", "1100"},
		{"110011111", "10011"},
		{"0111001010", "011000"},
		// Cadenas con un error
		{"0011111", "0011"},          // Cadena sin error: 0011110
		{"101001101", "11001"},       // Cadena sin error: 101001101
		{"100100001011", "10010000"}, // Cadena sin error: 100100001010
	}

	fails := []struct {
		input  string
		output string
	}{
		// Cadenas con dos errores, tienen que fallar
		{"0011101", "0011"},          // Cadena sin error: 0011110
		{"101001100", "11001"},       // Cadena sin error: 111001101
		{"101100001011", "10010000"}, // Cadena sin error: 100100001010
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			decoded := HammingDecode(test.input)
			if decoded != test.output {
				t.Errorf("For message %s, expected %s but got %s", test.input, test.output, decoded)
			}
		})
	}

	for _, test := range fails {
		t.Run(test.input, func(t *testing.T) {
			decoded := HammingDecode(test.input)
			if decoded == test.output {
				t.Errorf("For message %s, expected to fail but got %s", test.input, decoded)
			}
		})
	}
}
