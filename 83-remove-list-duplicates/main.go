/*
* Algo:
- head is a sorted singly-linked list
- Identify the first value in the linked list e.g. curr = head
- Iterate over the list from the second value e.g. for curr.Next != nil { work inside here }
- Apply evaluation condition whereby if the current value == next, evaluate the next next val
- e.g. if curr.Val == curr.Next.Val, set curr.Next = curr.Next.Next to evaluate the next value
- if curr.Next is evaluated to be a unique value set curr = curr.Next
- return head
*/

package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode{
	if head == nil {
        return nil
    }
    
    curr := head  // Start with the first node
    
    for curr.Next != nil {
        if curr.Val == curr.Next.Val {
            // Skip the duplicate node
            curr.Next = curr.Next.Next
        } else {
            // Move to the next distinct value
            curr = curr.Next
        }
    }
    
    return head
}