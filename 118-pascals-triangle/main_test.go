package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct{
		numRows 	int
		expected 	[][]int
	}{
		// Test case 1: numRows = 5, output [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
		{},

		// Test case 2: numRows = 1, output [[1]]
		{},

	}

	// Evaluate test cases
	for _, test := range tests {
		result := generate(test.numRows)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("generate(%v) = %v, want: %v", test.numRows, result, test.expected)
		}
	}
}