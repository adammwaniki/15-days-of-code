/*
* Algo:
- Early return true if s is " " i.e. a string with no characters(different from empty string)
- Strip the string of whitespaces and characters not in the alphabet
- Join the characters and convert to lowercase to be case insensitive
- A string is a slice of bytes
- Option 1: Iterate from both ends of the string and compare iterators
    - for i, j := 0, len(s)-1 ; i < j ; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return true
        } else {
            return false
        }
    }
- Option 2: split the string in half and compare the values e.g. using reflect.DeepEqual
    - high := len(s) - 1
    - low := 0
    - mid := int((high + low)/2) // force it to round down
    if !reflect.DeepEqual(s[low:mid], s[mid:high]) {
        return false
    } else {
        return true
    }
*/
package main

func isPalindrome(s string) bool {
    if s == " " {
        return true
    }

}

func main(){}