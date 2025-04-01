package main

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	// Helper function to create a linked list with a cycle
	createLinkedList := func(values []int, pos int) *ListNode {
		if len(values) == 0 {
			return nil
		}

		dummy := &ListNode{}
		cur := dummy
		var cycleNode *ListNode

		for i, val := range values {
			newNode := &ListNode{Val: val}
			cur.Next = newNode
			cur = newNode

			if i == pos {
				cycleNode = newNode
			}
		}

		// Create cycle if pos is valid
		if cycleNode != nil {
			cur.Next = cycleNode
		}

		return dummy.Next
	}

	tests := []struct {
		head     *ListNode
		expected bool
	}{
		// Test case 1: input head = [3, 2, 0, -4], pos = 1, output true
		{createLinkedList([]int{3, 2, 0, -4}, 1), true},

		// Test case 2: input head = [1, 2], pos = 0, output true
		{createLinkedList([]int{1, 2}, 0), true},

		// Test case 3: input head = [1], pos = -1, output false
		{createLinkedList([]int{1}, -1), false},
	}

	for i, test := range tests {
		result := hasCycle(test.head)
		if result != test.expected {
			t.Errorf("Test %d failed: hasCycle() = %v; want: %v", i+1, result, test.expected)
		}
	}
}
