package main

import (
	"reflect"
	"testing"
)

func TestAddBinary(t *testing.T) {
	// Define a slice of test cases
	tests := []struct{
		a 		 string
		b 		 string
		expexted string
	}{
		// Test case 1: input a "11", b "1", expected output "100"
		{"11", "1", "100"},

		// Test case 2: input a "1010", b "1011", expected output "10101"
		{"1010", "1011", "10101"},
	}

	// Iterate over the test cases
	for _, test := range tests {
		result := addBinary(test.a, test.b)

		// Return an error in case it fails
		if !reflect.DeepEqual(result, test.expexted){
			t.Errorf("addBinary(%v, %v) = %v; want: %v", test.a, test.b, result, test.expexted)
		}
	}
}