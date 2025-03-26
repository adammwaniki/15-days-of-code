/*
*Algo:
- Iterate over the grid so that we can target the slices in a nested loop
- Iterate over each slice and produce a joined slice 
    e.g using slices.Concat() or the variadic operator ...
- For mathematically impossible cases return -1,
    case1: evaluate for num > x { if num%x != 0 {return -1}}
    case2: evaluate for num < x { if x%num != 0 {return -1}}
- For the remaining cases, find the middle number (round down like binSearch)
- find the number of steps needed to make each num equal to mid
    steps := (num - mid)/x
- Return the sum of steps
*/

package main

import "slices"

func minOperations(grid [][]int, x int) int {
    // Handling egde cases where the grid is empty
    if len(grid) == 0 {
        return -1 
    }

    grandLine := grid[0]
    for _, gridLine := range grid[1:] {
        grandLine = append(grandLine, gridLine...)
    } 
    slices.Sort(grandLine)
    
    mid := grandLine[1]
    steps := 0
    for _, num := range grandLine {
        for num > x {
            ans := num % x
            if ans != 0 {
                return -1
            } else {
                steps = (num - mid)/x
                steps += steps
            }
        } 
        for num < x { 
            ans := x % num
            if ans != 0 {
                return -1
            } else {
                steps = (num - mid)/x
                steps += steps
            }
        }
    }
    return steps
    
}

func main(){}