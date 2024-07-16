package medium

import "sort"

// 1. Two Sum : Check if a pair with given sum exists in Array
// Problem Statement: Given an array of integers arr[] and an integer target.
// 1st variant: Return YES if there exist two numbers such that their sum is equal to the target.
// Otherwise, return NO.
// 2nd variant: Return indices of the two numbers such that their sum is equal to the target.
// Otherwise, we will return {-1, -1}.

// Note: You are not allowed to use the same element twice.
// Example: If the target is equal to 6 and num[1] = 3, then nums[1] + nums[1] = target is not a solution.

//Brute force
//Better - by making j := i+1 can ignore if i == j{continue}

// combined 1st and 2nd variant
func TwoSumBrute(arr []int, target int) (string, [2]int) {

	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == target {
				return "YES", [2]int{i, j}
			}

		}
	}
	return "NO", [2]int{-1, -1}
}

//Optimal (using hash maps)

func TwoSumOptimal(arr []int, target int) (string, [2]int) {
	numMap := make(map[int]int)

	for i, num := range arr {
		compliment := target - num
		if index, found := numMap[compliment]; found {
			return "YES", [2]int{index, i}
		}
		numMap[num] = i
	}

	return "NO", [2]int{-1, -1}

}

//optimal2 (without using hashmap) using two pointer
// this optimal only for 1st variant, not for 2nd as, it used lot of space and time complexity varies

// Pair struct to hold the value and the original index
type Pair struct {
	value int
	index int
}

func TwoSumOptimal2(arr []int, target int) (string, [2]int) {
	n := len(arr)
	// Create a slice of pairs to keep track of original indices
	pairs := make([]Pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = Pair{arr[i], i}
	}

	// Sort the pairs slice based on the values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	left := 0
	right := n - 1
	for left < right {
		sum := pairs[left].value + pairs[right].value
		if sum == target {
			return "YES", [2]int{pairs[left].index, pairs[right].index}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return "NO", [2]int{-1, -1}
}

//Test case
// arr1 := []int{2, 6, 5, 8, 11}
// target1 := 14

// result1, indices1 := medium.TwoSumOptimal2(arr1, target1)
// fmt.Printf("Result: %s, Indices: %v\n", result1, indices1) // Output: Result: YES, Indices: [1, 3]

// arr2 := []int{2, 6, 5, 8, 11}
// target2 := 15

// result2, indices2 := medium.TwoSumOptimal2(arr2, target2)
// fmt.Printf("Result: %s, Indices: %v\n", result2, indices2) // Output: Result: NO, Indices: [-1, -1]

//----------------------------------------------------------------------------------------------------

// 2. Sort an array of 0s, 1s and 2s

// Problem Statement: Given an array consisting of only 0s, 1s, and 2s.
// Write a program to in-place sort the array without using inbuilt sort functions.
// ( Expected: Single pass-O(N) and constant space)

func SortArrZeroOneAndTwosBrute(arr []int) []int {

	return arr
}
