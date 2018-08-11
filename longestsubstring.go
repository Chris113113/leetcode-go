package main

import ("log")

func main() {
    log.Println(lengthOfLongestSubstring("shadawadeboob"))
}

func lengthOfLongestSubstring(s string) int {
	max, start, end := 0, 0, 0
	substr := ""

	for ; end < len(s); end++ {
		substr = s[start:end]

	}

	return max
}