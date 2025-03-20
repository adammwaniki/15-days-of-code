package main

func plusOne(digits []int) []int {
    var target int = digits[len(digits)-1]
	    
	if target == 9 {
		digits[len(digits)-1] = 1
		digits = append(digits, 0)
	} else {
		digits[len(digits)-1] = target + 1
	}
    return digits
}

func main (){}