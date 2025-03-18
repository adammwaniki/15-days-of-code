package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	// Define a slice of test cases
	tests := []struct {
		nums 	[] int 	// The input slice of numbers
		target 	int 	// The target number 
		expected 	int		// The expected index of the target number 
	}{
		// Test case 1: target 5 from [1,3,5,6], expected output 2
		{[]int{1,3,5,6}, 5, 2},

		// Test case 2: target 2 from [1,3,5,6], expected output 1
		{[]int{1,3,5,6}, 2, 1},

		// Test case 3: target 7 from [1,3,5,6], expected output 4
		{[]int{1,3,5,6}, 7, 4},
	}

	// Iterate over each test case
	for _, test := range tests {
		// Call the searchInsert function with the test case inputs
		result := searchInsert(test.nums, test.target)

		// Check if the result matches the expected output
		if result != test.expected {
			// If the test fails, report an error
			t.Errorf("searchInsert(%v, %d) = %d; want %d", test.nums, test.target, result, test.expected)
			// t.Errorf prints the error without stopping the test execution
		}
	}
}