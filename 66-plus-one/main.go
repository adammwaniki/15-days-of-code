package main

func plusOne(digits []int) []int {
	// A variable carry is initialized to 1 to represent the +1 increment.
    carry := 1  // Start with +1

	// The loop iterates backward from the last digit (rightmost).
	// The last digit is increased by 1 (adding carry).
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i] += carry
        if digits[i] < 10 {
            return digits  // No carry needed, return early
        }
        digits[i] = 0  // Carry over, set current digit to 0
    }

	// Handling the Edge Case Where All Digits Are 9
    // If loop completes, it means we had a carry that was not absorbed (e.g., 999 → 1000)
    return append([]int{1}, digits...)
}

func main (){}