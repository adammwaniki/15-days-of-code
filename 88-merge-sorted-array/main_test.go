package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1		[]int
		nums2		[]int
		m			int
		n			int
		expected	[]int

	}{
		// Test case 1: input nums1 = [1,2,3,0,0,0], m = 3, nums2[2,5,6], n = 3, output [1,2,2,3,5,6] 
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},

		// Test case 2: input nums1 = [1], m = 1, nums2 = [], n = 0, output [1]
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		
		// Test case 3: input nums1 = [0], m = 0, nums2 = [1], n = 1, output [1]
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, test := range tests {
		// Create a copy of nums1 to avoid modifying the original test cases
		nums1Copy := append([]int{}, test.nums1...)

		// Call the function
		// merge modifies nums1 in-place, so we copy nums1 before calling the function to avoid modifying the test case struct itself
		merge(nums1Copy, test.m, test.nums2, test.n)

		// Validate the result
		if !reflect.DeepEqual(nums1Copy, test.expected) {
			t.Errorf("merge(%v, %d, %v, %d) = %v; want %v",
				test.nums1, test.m, test.nums2, test.n, nums1Copy, test.expected)
		}
	}
}