package main

import "math/big"

func addBinary(a string, b string) string {
    // Initialise a zero-value big.Int and allow setting its value later
	num1 := new(big.Int)
	num2 := new(big.Int)
    // alternatively we could have used num := big.NewInt(value) but the value is limited to numbers in the int64 range

	// Parse binary strings into big.Int values
	num1.SetString(a, 2)
	num2.SetString(b, 2)

	// Perform binary addition
	sum := new(big.Int)
	sum.Add(num1, num2)

	// Convert the result back to a binary string
	return sum.Text(2)
}

func main() {
    // main_test.go will handle the unit tests
}
