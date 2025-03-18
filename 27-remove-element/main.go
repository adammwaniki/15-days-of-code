package main

import (
	"slices"
)

func removeElement(nums []int, val int) int {
    // DeleteFunc removes any elements from a slice s for which the del func returns true, returning the modified slice
    nums = slices.DeleteFunc(nums, func(n int) bool{
        return n == val //delete the elements that are equal to val
    })
    count := len(nums)
    return count
}

func main() {} // We'll handle unit tests in the test file