/*
* Algo:
- Early return true if s is " " i.e. a string with no characters(different from empty string)
- Strip the string of whitespaces and characters not in the alphabet
- Join the characters and convert to lowercase to be case insensitive
    - This can be handled using the regexp package
- A string is a slice of bytes
- Option 1: Iterate from both ends of the string and compare iterators
  - i, j  := 0, len(matches)-1
    for i < j {
        if matches[i] != matches[j] {
            return false
        }
        i++
        j--
    }
    O(n) time

*/
package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
    // Early return true if length of string is 1
    if len([]rune(s)) == 1 {
        return true
    } 

    pattern := regexp.MustCompile(`[a-zA-Z0-9]`) // regex literal for the range of alpha-numeric characters instead of regex class \w to prevent underscore matching
    matches := pattern.FindAllString(strings.ToLower(s), -1)

    i, j  := 0, len(matches)-1
    for i < j {
        if matches[i] != matches[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func main(){}
