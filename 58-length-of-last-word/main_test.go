package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	// Define a slice of test cases
	tests := []struct{
		s 			string
		expected 	int
	}{
		// Test case 1: s is "Hello World", expected output is 5
		{"Hello World", 5},

		// Test case 2: s is "   fly me   to   the moon  ", expected output is 4
		{"   fly me   to   the moon  ", 4},

		// Test case 3: s is "luffy is still joyboy", expected output is 6
		{"luffy is still joyboy", 6},
	}

	// Iterating over each test case
	for _, test := range tests {
		// Call the lengthOfLastWord
		result := lengthOfLastWord(test.s)

		// Check if result matches the expected output
		if result != test.expected {
			// Report an error if the test fails
			// If the test fails, report an error
			t.Errorf("lengthOfLastWord(%v) = %d; want %d", test.s, result, test.expected)
			// t.Errorf prints the error without stopping the test execution
		}
	}
}