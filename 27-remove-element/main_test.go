package main  // The test file belongs to the same package as the main application

import (
	"testing" // Import the testing package for writing test cases
)

// TestRemoveElement is a test function that verifies the correctness of the removeElement function.
func TestRemoveElement(t *testing.T) {
	// Define a slice of test cases using a struct
	tests := []struct {
		nums     []int // The input slice of numbers
		val      int   // The value to be removed from the slice
		expected int   // The expected length of the slice after removing val
	}{
		// Test case 1: Removing 3 from [3, 2, 2, 3] should result in [2, 2] with length 2
		{[]int{3, 2, 2, 3}, 3, 2},

		// Test case 2: Removing 2 from [0, 1, 2, 2, 3, 0, 4, 2] should result in [0, 1, 3, 0, 4] with length 5
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},

		// Test case 3: Removing 1 from [1, 1, 1, 1] should result in an empty slice with length 0
		{[]int{1, 1, 1, 1}, 1, 0},

		// Test case 4: Removing 6 from [1, 2, 3, 4, 5] (which doesn’t exist in the slice) should result in the same length 5
		{[]int{1, 2, 3, 4, 5}, 6, 5},
	}

	// Iterate over each test case
	for _, test := range tests {
		// Call the removeElement function with the test case inputs
		result := removeElement(test.nums, test.val)

		// Check if the result matches the expected output
		if result != test.expected {
			// If the test fails, report an error
			t.Errorf("removeElement(%v, %d) = %d; want %d", test.nums, test.val, result, test.expected)
			// t.Errorf prints the error without stopping the test execution
		}
	}
}
