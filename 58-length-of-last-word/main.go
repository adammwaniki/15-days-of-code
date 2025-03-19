package main

import "strings"

func lengthOfLastWord(s string) int {
	// Use the Fields function from the strings package to separate
	// strings by whitespace and return a slice of strings
    sFields := strings.Fields(s)
	// Directly return the length of the element at the last index
    return len(sFields[len(sFields)-1])
}

func main () {}