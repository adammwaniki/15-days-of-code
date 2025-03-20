package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	// Define a slice of test cases
	tests := []struct{
		digits	 []int
		expected []int
	}{
		// Test case 1: input [1,2,3], expected output [1,2,4]
		{[]int{1,2,3}, []int{1,2,4}},

		// Test case 2: input [4,3,2,1], expected output [4,3,2,2]
		{[]int{4,3,2,1}, []int{4,3,2,2}},

		// Test case 3: input [9], expected output [1,0]
		{[]int{9}, []int{1,0}},

		// Test case 4: input [9,8,9], expected output [9,9,0]
		{[]int{9,8,9}, []int{9,9,0}},
	}

	// Iterate over the test cases
	for _, test := range tests {
		result := plusOne(test.digits)

		// Check if results match the expected output
		// Using deep reflect because arrays can only use the comparison operator someSlice == nil
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("plusOne(%#v) = %#v; want %#v", test.digits, result, test.expected)
		}

		// Alternatively I could use the slices.Equal function from the slices package
		/*
		if !slices.Equal(result, test.expected) {
			t.Errorf("plusOne(%v) = %v; want %v", test.digits, result, test.expected)
		}
		*/
	}
}


