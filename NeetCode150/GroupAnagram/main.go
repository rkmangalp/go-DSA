package main

import "fmt"

// 49. Group Anagrams
// Medium

// Given an array of strings strs, group the anagrams together.
// You can return the answer in any order.
// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Explanation:
// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

// Example 2:
// Input: strs = [""]
// Output: [[""]]

// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]

func main() {

	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, s := range strs {
		var freq [26]int
		for _, ch := range s {
			freq[ch-'a']++
		}
		m[freq] = append(m[freq], s)
	}

	res := make([][]string, 0, len(m))
	for _, group := range m {
		res = append(res, group)
	}

	return res
}

// m := make(map[[26]int][]string)

// for _, s := range strs {
// 	var freq [26]int
// 	for _, ch := range s {
// 		freq[ch-'a']++
// 	}
// 	m[freq] = append(m[freq], s)
// }

// res := make([][]string, 0, len(m))
// for _, group := range m {
// 	res = append(res, group)
// }

// return res
