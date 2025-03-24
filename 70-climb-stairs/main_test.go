package main

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct{
		n int
		expected int
	}{
		// Test case: 1, input 2, expected output 2 
		{2, 2},

		// Test case: 2, input 3, expected output 3
		{3, 3},

		// Test case: 3, input 4 expected output 5
		{4, 5},

		// Test case: 4, input 13 expected output 377
		{13, 377},
	}

	for _, test := range tests {
		result := climbStairs(test.n)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("climbStairs(%d) = %d, want: %d", test.n, result, test.expected)
		}
	}
}