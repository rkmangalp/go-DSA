package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return true if any
// value appears at least twice in the array,
// and false if every element is distinct.
// Examples
// Example 1:
// Input: nums = [1,2,3,1]
// Output: true (The element 1 occurs twice)
// Example 2:
// Input: nums = [1,2,3,4]
// Output: false (All elements are distinct)

func main() {
	arr := []int{1, 2, 3, 2}
	res := containDupOptimal(arr)
	fmt.Println(res)
}

func containDupOptimal(arr []int) bool {
	hashset := make(map[int]struct{})

	for _, num := range arr {
		if _, ok := hashset[num]; ok {
			return true
		}
		hashset[num] = struct{}{}
	}

	return false
}

func containDupBetter(arr []int) bool {
	sort.Ints(arr)
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}

func containDupBrute(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}
