/*
* Algo
- The values in a particular row n of Pascal's triangle can be found multiplicatively
- C(n, k) = C(n, k-1) * (n - k + 1)/(k)
- That is:
- row[k] = row[k-1] * int(n - k +1)/int(k)
- O(n) time, O(n) space
*/

package main

func getRow(rowIndex int) []int {
    row := make([]int, rowIndex+1) // the rows in Pascal's triangle have a length of n+1
    row[0] = 1 // the first value is always 1

    for k := 1 ; k <= rowIndex ; k++ {
        row[k] = row[k-1] * int(rowIndex - k +1)/int(k)
    }
    
    return row
}

func main(){}