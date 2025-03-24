/*
Algo logic:
Binary Addition Rules

XOR (^) acts like normal addition without carrying:

0 ^ 0 = 0
0 ^ 1 = 1
1 ^ 0 = 1
1 ^ 1 = 0 (ignores carry)

AND (&) + left shift (<<) calculates the carry:

1 & 1 = 1 (carry occurs)

The carry is then shifted left (<< 1) to add it to the next higher bit.
*/
package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	// Convert binary strings to integers
    // This is how strconv.ParseInt() works:
    // func ParseInt(s string, base int, bitSize int) (i int64, err error)
	num1, err1 := strconv.ParseInt(a, 2, 64)
    // Since a is a binary number represented as a string,
    // we pass 2 as the base to indicate that it's in binary format.
    // Bit size 64 tells Go that the result should fit in an int64 type,
    // meaning it can store large values.
	if err1 != nil {
		fmt.Println("Error parsing a:", err1)
		return ""
	}

	num2, err2 := strconv.ParseInt(b, 2, 64)
	if err2 != nil {
		fmt.Println("Error parsing b:", err2)
		return ""
	}

	// Perform binary addition using bitwise operations
	for num2 != 0 {
		sum := num1 ^ num2        // XOR for sum
		carry := (num1 & num2) << 1 // AND finds carry and shifts left
		num1 = sum                // Store sum
		num2 = carry              // Update carry for next iteration
	}

	// Convert the result back to a binary string
	return strconv.FormatInt(num1, 2)
}

func main() {
    // main_test.go will handle the unit tests
}
