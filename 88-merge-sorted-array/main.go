/*
* Algo:
- The naive approach is to concatenate the arrays and then sort the resulting array
- Time complexity: O((m + n) log(m + n))
- Space complexity: O(1)
- Ideal approach is merge sort
- Time complexity: Merge Sort O(m + n)
- Space complexity: O(n1 + n2)
    Typical merge sort
    - Create an array to store the sorted values arr3
    - Simultaneously traverse arr1 and arr2 from left to right
    - For each current element in arr1 and arr2, pick the smaller one and copy it to arr3
    - Move a step ahead in the array that was picked and evaluate and copy again
    - If there are any elements left in arr1 or arr2, append them to arr3
- However, in this case we need to handle it in place, into the extra space allocated in nums1
- Space complexity O(1)
- As such we shall slightly change how we traverse the arrays
    - Traverse the arrays from right to left, that is largest to smallest value
        i.e. the elements corresponding to index m-1 and n-1
    - Instead of copying as before, we shall place the values into arr1
    - Placement will prevent shifting of values
    - Placement will start from the very end of arr1 i.e. 
        at the index corrresponding to (m+n)-1
*/
package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
    // Early return for edge cases where nums2 is empty
    if n == 0 {
        return 
    }
    
    // Start merging from the end of nums1 and nums2
    i, j, k := m-1, n-1, m+n-1
    
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
    
    // If any elements left in nums2, copy them
    for j >= 0 {
        nums1[k] = nums2[j]
        j--
        k--
    }
}

func main (){}