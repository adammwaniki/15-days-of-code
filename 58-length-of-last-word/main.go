package main

import "strings"

func lengthOfLastWord(s string) int {
    sFields := strings.Fields(s)
    ans := 0
    for i, v := range sFields {
        if i == len(sFields)-1 {
            ans += len(v)
        }
    }
    return ans
}

func main () {}