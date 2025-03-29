/*
*Algo:
- Pascal's triangle serves to find the binomial coefficients of terms of binomial expression (x + y)^n
- The values in a particular row of the triangle can be found by evaluating C(n, k)
- Mathematically this will be evaluated as: C(n, k) = (n!)/(k!(n - k)!)
- Factorials can be very large numbers hence for programming this can be optimised
- We can apply the multiplicative approach. 
- This will allow us to generate each value iteratively
- C(n, k) = C(n, k-1) * (n - k + 1)/(k)
*
- If our objective was to only find the values in a particular row n, the following would suffice:
    // generate computes Pascal's Triangle row n
    func generate(n int) []int {
        row := make([]int, n+1)
        row[0] = 1 // First element is always 1

        for k := 1; k <= n; k++ {
            row[k] = row[k-1] * int(n-k+1) / int(k)
        }

        return row
    }
- O(n) Time, O(n) Space
*
- However,  we also need to include the values from the rows prior to the current
- 
*/

package main

func generate(numRows int) [][]int {
    triangle := make([][]int, numRows) // Create a 2D slice

	for i := 0; i < numRows; i++ {
		// Create a new row with size i+1
		row := make([]int, i+1)
		row[0], row[i] = 1, 1 // First and last element are always 1

		// Compute values using the iterative formula
		for k := 1; k < i; k++ {
			row[k] = triangle[i-1][k-1] + triangle[i-1][k] // Sum of the two elements above
		}

		// Append the computed row to the triangle
		triangle[i] = row
	}

	return triangle
}
func main() {}