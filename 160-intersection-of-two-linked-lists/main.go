/*
* Algo:
- Intersection of two lists typically means that they may be of different lengths until the point of intersection
- After intersecting, the length remains the same for both until the end
- We can find the point of intersection by starting comparison evaluation from the first node where
	it takes the same number of steps on each list to reach the intersection
- i.e. Find the difference in the length of the lists then set the starting point to be the corresponding
	node subtracted from the longer list

* Pseudocode:
- Helper function to evaluate the length of the lists
	- helperFunc(n *ListNode) (result int) {
		if node n is not nil {
			set n = n.Next
			recursively return 1 + helperFunc(n)
		}
		return 0
	}

- getIntersectionNode(headA, headB *ListNode) *ListNode {
	- Find the difference in length between headA and headB
	- Initialise two dummy pointers to traverse headA and headB 
		e.g., p := headA, q := headB 

	- Set the starting points for list traversal
	if diff > 0 {
		for i := 0 ; i < diff ; i++ {
			p = p.Next
		}
	} else {
		for i := 0 ; i < -diff ; i++ {
			q = q.Next
		} 
	}
	
	- Begin list traversal while comparing p to q
	- NB: evaluating if *p == *p will compare the value held by p and q not the actual location in memory
	for {
		if p == nil || p.Next == nil {
			return nil 
		}
		if p == q {
			return p
		}
		// move p and q to the next values if no match
		p = p.Next
		q = q.Next
	}
	return nil

  }
*/

package main

type ListNode struct {
	Val int
	Next *ListNode
}

// Helper function to find the length of a linked list
func recursiveLength(n *ListNode) int {
    if n == nil {
        return 0
    }
    return 1 + recursiveLength(n.Next)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    lenA := recursiveLength(headA)
    lenB := recursiveLength(headB)
    diff := lenA - lenB

    p, q := headA, headB

    // Align p and q to start at the same position
    if diff > 0 {
        for i := 0; i < diff; i++ {
            p = p.Next
        }
    } else {
        for i := 0; i < -diff; i++ {
            q = q.Next
        }
    }

    // Traverse together to find intersection
    for p != nil && q != nil {
        if p == q {
            return p
        }
        p = p.Next
        q = q.Next
    }

    return nil // No intersection found
}

func main() {}