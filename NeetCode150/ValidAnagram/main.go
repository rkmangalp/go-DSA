package main

import "fmt"

// Given two strings s and t, return true if t is
// an anagram of s, and false otherwise.

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagramOptimal(s, t))
	fmt.Println(isAnagramMap(s, t))
}

func isAnagramOptimal(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}

func isAnagramMap(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}
