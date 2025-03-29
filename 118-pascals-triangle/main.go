package main

// generate computes Pascal's Triangle row n
func generate(n int) []int {
	row := make([]int, n+1)
	row[0] = 1 // First element is always 1

	for k := 1; k <= n; k++ {
		row[k] = row[k-1] * int(n-k+1) / int(k)
	}

	return row
}

func main() {}