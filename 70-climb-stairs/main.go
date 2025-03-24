// This is an implementation of the fibonacci sequence
// We can solve this using Binet's formula for optimised efficiency with large numbers
// F(n) = ((((1 + √5) / 2) ^ n) - (((1 - √5) / 2) ^ n)) / √5
// In the above formula: 
//      (1 + √5) / 2 is the golden ratio often denoted as φ or phi
//      (1 - √5) / 2 is the negative reciprocal of the golden ratio
// This can be translated as:
// F := int((math.Pow(1+math.Sqrt(5), float64(n)) - math.Pow(1-math.Sqrt(5), float64(n))) / math.Sqrt(5))
// O(1) Time  

package main

import "math"

func climbStairs(n int) int {
    sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	psi := (1 - sqrt5) / 2
	// Round the result to the nearest integer to avoid precision issues
	return int(math.Round((math.Pow(phi, float64(n+1)) - math.Pow(psi, float64(n+1))) / sqrt5))
}

func main (){}