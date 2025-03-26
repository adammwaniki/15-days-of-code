package main

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice (for easier comparison)
func linkedListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		head     *ListNode
		expected *ListNode
	}{
		// Test case 1: input [1,1,2], output [1,2]
		{createLinkedList([]int{1, 1, 2}), createLinkedList([]int{1, 2})},

		// Test case 2: input [1,1,2,3,3], output [1,2,3]
		{createLinkedList([]int{1, 1, 2, 3, 3}), createLinkedList([]int{1, 2, 3})},
	}

	for _, test := range tests {
		result := deleteDuplicates(test.head)
		if !reflect.DeepEqual(linkedListToSlice(result), linkedListToSlice(test.expected)) {
			t.Errorf("deleteDuplicates(%v) = %v, want: %v", linkedListToSlice(test.head), linkedListToSlice(result), linkedListToSlice(test.expected))
		}
	}
}
