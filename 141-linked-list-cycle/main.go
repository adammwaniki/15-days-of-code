package main

/*
* Algo:
 - Apply Floyd's Tortoise and Hare Algorithm or Brent's algorithm
 - Option 1: Tortoise and Hare
 - The algorithm maintains two pointers into the given sequence
 - One (the tortoise) at x_i, and the other (the hare) at x_2i
 - At each step of the algorithm, it moves i forward by 1
 - This results in the tortoise moving forward by 1 step and the hare by 2 steps
 - It then compares the sequence values at these two pointers
 - The smallest value of i > 0 for which the tortoise and hare point to equal values is the desired value ν

 * Pseudocode: 
 - Declare and initialise tortoise and hare as the head so that they can start traversing the list together
 - While neither hare nor hare.Next is nil: tortoise = tortoise.Next, hare = hare.Next.Next
 - If tortoise == hare, break the while loop // In this case return true
 - This means that a cycle has been found
 - If tortoise != hare return false

 *Bonus: Identify the beginning of the cycle
 - In order to identify the node where the cycle begins
 - From where we broke the loop // if we used break instead of return
 - Send the tortoise back to the start and make the hare and rabbit move one step at a time each
 - Where they meet is the node that starts the cycle
 - Set the tortoise = head
 - While tortoise != hare: tortoise = tortoise.Next, hare = hare.Next
 - Return tortoise or hare since they will be at the same place

 Time: O(n), Space: O(1)

*/
type ListNode struct {
	Val int
	Next *ListNode
}



func hasCycle(head *ListNode) bool {
	// Declare tortoise and hare for list traversal
    tortoise := head
	hare := head

    for { // for loop with absent condition is evaluated true
		if hare == nil || hare.Next == nil {
			return false
		}

        tortoise = tortoise.Next
		hare = hare.Next.Next

		if tortoise == hare {
			return true
		}
    }
    
}