package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{7, "7"},
		{30, "FizzBuzz"},
		{9, "Fizz"},
		{10, "Buzz"},
	}

	for _, tc := range tests {
		result := fizzbuzz(tc.input)
		if result != tc.expected {
			t.Errorf("fizzbuzz(%d) = %s, want %s", tc.input, result, tc.expected)
		}
	}
}
