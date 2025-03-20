package main

import "strconv"

func plusOne(digits []int) []int {
    var ans int = (digits[len(digits)-1] + 1)
    
	if ans > 9 {
		numStr := strconv.Itoa(ans)
		for _, v := range numStr {
			// Since int(v) directly converts a rune (character) to its ASCII value
			// subtract '0' to get the integer value
			digits = append(digits, int(v - '0'))
		}
	} else {
		digits[len(digits)-1] = ans
	}
    return digits
}

func main (){}