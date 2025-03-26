/*
*Algo:
- Iterate over the grid so that we can target the slices in a nested loop
- Iterate over each slice and produce a joined slice e.g using slices.Concat() 
- For mathematically impossible cases return -1,
    case1: evaluate for num > x { if num%x != 0 {return -1}}
    case2: evaluate for num < x { if x%num != 0 {return -1}}
- For the remaining cases, find the middle number (round down like binSearch)
- find the number of steps needed to make each num equal to mid
    steps := (num - mid)/x
- Return the sum of steps
*/

package main

func minOperations(grid [][]int, x int) int {
    
}

func main(){}