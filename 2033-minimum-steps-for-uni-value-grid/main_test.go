package main

import (
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T){
	tests := []struct{
		grid [][]int
		x int
		expected int
	}{
		// Test case 1: input [[2,4],[6,8]], x is 2, output 4
		{[][]int{
			{2,4},
			{6,8},
			},
			2, 4},
		
		// Test case 2: input [[1,5],[2,3]], x is 1, output 5
		{[][]int{{1,5}, {2,3}}, 1, 5},

		// Test case 3: input [[1,2],[3,4]], x is 2, output -1
		{[][]int{{1,2}, {3,4}}, 2, -1},
	}
	for _, test := range tests {
		result := minOperations(test.grid, test.x)
		if !reflect.DeepEqual(test.expected, result){
			t.Errorf("minOperations(%v, %v) = %v, expected %v", test.grid, test.x, result, test.expected)
		}
	}
}