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

import "strconv"

func addBinary(a string, b string) string {
    a = strconv.Atoi(a)
    b = strconv.Atoi(b)
    for b != 0 { 
        sum := a ^ b      // XOR for addition without carry
        carry := (a & b) << 1 // AND finds carry and shifts it left
        a = sum           // Store sum
        b = carry         // Update carry for next iteration
    }
    return a
}

func main(){}