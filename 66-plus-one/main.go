package main

func plusOne(digits []int) []int {
    var ans int = (digits[len(digits)-1] + 1)
    digits[len(digits)-1] = ans
    return digits
}

func main (){}