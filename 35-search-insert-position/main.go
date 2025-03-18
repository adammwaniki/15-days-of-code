package main

func searchInsert(nums []int, target int) int {
    // Start by implementing simple binary search for O(log n) complexity
    high := len(nums) - 1
    low := 0
    for low <= high {
        mid := (high + low)/2
        guess := nums[mid]
        if guess == target{
            return mid
        }
        if guess > target{
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return low
}

func main(){} // Unit tests are handled in main_test.go