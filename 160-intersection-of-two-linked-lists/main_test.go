package main

import "testing"

// Helper function to create linked lists for the test cases
// Function takes in a slice of values and returns a pointer to ListNode
func createLinkedList(values []int) *ListNode {
	// Handle edge cases where there are no values
	if len(values) == 0 {
		return nil
	}
	// Initialise a linked list head
	head := &ListNode{Val: values[0]}
	
	// Initialise a dummy pointer to the linked list that will allow us to traverse the slice values
	curr := head

	// Loop over the values in the range of the slice
	// Set the curr.Next = the value in the range of values i.e. curr.Next = &ListNode{Val: value}
	// Set curr = curr.Next to allow progress to the next step in the iteration
	// Return the head

	for _, value := range values[1:] { // start from the second element since we've already set the first
		curr.Next = &ListNode{Val: value}
		curr = curr.Next
	}
	return head
}

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		headA		*ListNode
		headB		*ListNode
		expected	*ListNode
	}{
		// Test case 1: input headA = [4,1,8,4,5] headB = [5,6,1,8,4,5], output expect intersection at '8'
		{}, 
	}

	for _, test := range tests {
		result := getIntersectionNode(test.headA, test.headB)

	}
}