package main

// generatePascalsRow computes Pascal's Triangle row n
func generate(n int) []int64 {
	row := make([]int64, n+1)
	row[0] = 1 // First element is always 1

	for k := 1; k <= n; k++ {
		row[k] = row[k-1] * int64(n-k+1) / int64(k)
	}

	return row
}

func main() {}