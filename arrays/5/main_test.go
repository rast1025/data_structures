package main

import (
	"fmt"
	"testing"
)

func TestCompression(t *testing.T) {
	tt := []struct{
		input string
		expected string
	}{
		{"aaaabccccdde","a4b1c4d2e1"},
		{"aabbc", "aabbc"},
		{"adddddceee", "a1d5c1e3"},
	}

	for _, test := range tt {
		testName := fmt.Sprintf("input: %s, expected: %s", test.input, test.expected)
		t.Run(testName, func(t *testing.T){
			result := compress(test.input)
			if result != test.expected {
				t.Errorf("expected %s, got %s\n", test.expected, result)
			}
		})
	}
}
