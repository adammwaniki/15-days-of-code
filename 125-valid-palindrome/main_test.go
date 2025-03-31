package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s 			string
		expected 	bool
	}{
		// Test case 1: input s = "A man, a plan, a canal: Panama" , output true
		{"A man, a plan, a canal: Panama", true},

		// Test case 2: input s = "race a car", output false
		{"race a car", false},

		// Test case 3: s = " ", output true
		{" ", true},
	}

	for _, test := range tests {
		result := isPalindrome(test.s)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("isPalindrome(%v) = %v, want: %v", test.s, result, test.expected)
		}
	}
}