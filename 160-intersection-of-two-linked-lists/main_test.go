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

// Helper function to create an intersection
func createIntersectingLists(aVals, bVals []int, intersectVal int) (*ListNode, *ListNode, *ListNode) {
	headA := createLinkedList(aVals)
	headB := createLinkedList(bVals)

	// Find the intersection node in headA
	var intersectionNode *ListNode
	for node := headA; node != nil; node = node.Next {
		if node.Val == intersectVal {
			intersectionNode = node
			break
		}
	}

	// Attach headB to the intersection node
	if intersectionNode != nil {
		curr := headB
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = intersectionNode
	}

	return headA, headB, intersectionNode
}


func TestGetIntersectionNode(t *testing.T) {
	// Create the first test case with intersection
	headA1, headB1, expected1 := createIntersectingLists([]int{4, 1, 8, 4, 5}, []int{5, 6, 1}, 8)

	tests := []struct {
		headA, headB *ListNode
		expected     *ListNode
	}{
		// Test case 1: Lists intersect at node with value 8
		{headA1, headB1, expected1},

		// Test case 2: No intersection
		{
			createLinkedList([]int{2, 6, 4}),
			createLinkedList([]int{1, 5}),
			nil,
		},
	}

	for i, test := range tests {
		result := getIntersectionNode(test.headA, test.headB)

		if result != test.expected {
			t.Errorf("Test %d failed: got %v, expected %v", i+1, result, test.expected)
		}
	}
}

