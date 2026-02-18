package main

import (
	"fmt"
	"sort"
)

// Given an array of integers nums and an integer target, return indices of the two
// numbers such that they add up to target.

// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

func main() {

	arr1 := []int{2, 4, 7, 11, 15, 6}
	fmt.Println(twoSum(arr1, 9))

}

func twoSum(arr []int, tar int) []int {
	arrMap := make(map[int]int)
	for i, num := range arr {
		if idx, ok := arrMap[tar-arr[i]]; ok {
			return []int{i, idx}
		}
		arrMap[num] = i
	}
	return nil
}

// 💡 Intuition:
// Check every pair of numbers. If the sum equals target,
// return their indices.
func twoSumBrute(arr []int, tar int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == tar {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Intuition:
// Sort the array.
// Use two pointers: one at start, one at end.
// Move pointers inward depending on sum.
func twoSumBetter(arr []int, tar int) []int {

	n := len(arr)
	type pair struct {
		val  int
		indx int
	}
	sortArr := make([]pair, n)
	for i, v := range arr {
		sortArr[i] = pair{v, i}
	}
	sort.Slice(sortArr, func(i, j int) bool {
		return sortArr[i].val < sortArr[j].val
	})

	start := 0
	end := len(arr) - 1
	for start < end {
		sum := sortArr[start].val + sortArr[end].val
		if sum == tar {
			return []int{sortArr[start].indx, sortArr[end].indx}
		} else if sum < tar {
			start++
		} else {
			end--
		}
	}
	return nil
}

// Input: nums = [2,7,11,15], target = 9
func twoSumOptimal(arr []int, tar int) []int {

	arrMap := make(map[int]int)
	for i, num := range arr {
		if idx, ok := arrMap[tar-arr[i]]; ok {
			return []int{idx, i}
		}
		arrMap[num] = i
	}

	return nil
}
