package main

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	// Initialise a slice of test case structs to evaluate
	tests := []struct {
		rowIndex int
		expected []int
	}{
		// Test case 1: input rowIndex = 3, output [1,3,3,1]
		{3, []int{1,3,3,1}},

		// Test case 2: input rowIndex = 0, output [1]
		{0, []int{1}},

		// Test case 3: input rowIndex = 1, output [1,1]
		{1, []int{1,1}},

		// Test case 4: input rowIndex = 7, output [1,7,21,35,35,21,7,1]
		{7, []int{1,7,21,35,35,21,7,1}},

	}

	// Evaluate test cases
	for _, test := range tests {
		// The computed result
		result := getRow(test.rowIndex)

		// Compare the computed result against the expected output and return error if conflicts
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("getRow(%v) = %v, want: %v", test.rowIndex, result, test.expected)
		}
	}

	
}