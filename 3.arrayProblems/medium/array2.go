package medium

import (
	"fmt"
	"sort"
)

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

//Brute using Merge sort

func SortArrZeroOneAndTwosBrute(arr []int, low, high int) {

	if low >= high {
		return
	}
	mid := (low + high) / 2

	SortArrZeroOneAndTwosBrute(arr, low, mid)
	SortArrZeroOneAndTwosBrute(arr, mid+1, high)
	merge(arr, low, mid, high)

}
func merge(arr []int, low int, mid int, high int) {
	temp := make([]int, 0, high-low+1) // Temporary slice to hold the merged elements
	left := low                        // Starting index for the left subarray
	right := mid + 1                   // Starting index for the right subarray

	// Merge the two subarrays into temp
	for left <= mid && right <= high {
		if arr[left] <= arr[right] { // Compare elements from both subarrays
			temp = append(temp, arr[left]) // Add the smaller element to temp
			left++                         // Move to the next element in the left subarray
		} else {
			temp = append(temp, arr[right]) // Add the smaller element to temp
			right++                         // Move to the next element in the right subarray
		}
	}

	// Append the remaining elements of the left subarray, if any
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	// Append the remaining elements of the right subarray, if any
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	// Copy the sorted elements back into the original array
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

// Better approach

func SortArrZeroOneAndTwosBetter(arr []int) []int {
	var count_0, count_1, count_2 int
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			count_0++
		} else if arr[i] == 1 {
			count_1++
		} else {
			count_2++
		}
	}

	for i := 0; i < count_0; i++ {
		arr[i] = 0
	}
	for i := count_0; i < count_0+count_1; i++ {
		arr[i] = 1
	}
	for i := count_1 + count_0; i < n; i++ {
		arr[i] = 2
	}
	return arr
}

// {1, 2, 1, 2, 1, 2, 0, 0, 1, 1},
// Optimal approach - Dutch National flag algorithm
func SortArrZeroOneAndTwosOptimal(arr []int) {
	n := len(arr)
	low := 0
	mid := 0
	high := n - 1

	for mid <= high {
		switch arr[mid] {
		case 0:
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		default:
			fmt.Println("The number is not 0 or 1 or 2")
		}
	}
}

//Test case
// testCases := [][]int{
// 	{0, 1, 2, 0, 1, 2},
// 	{2, 2, 1, 0, 1, 0},
// 	{2, 1, 0},
// 	{1, 0, 2, 1, 0, 2},
// 	{1, 2, 1, 2, 1, 2},
// 	{2, 0, 0, 1, 2, 1},
// 	{1, 2, 1, 2, 1, 2, 0, 0, 1, 1},
// }

// for _, testCase := range testCases {
// 	fmt.Printf("Original array: %v\n", testCase)
// 	medium.SortArrZeroOneAndTwosOptimal(testCase)
// 	fmt.Printf("Sorted array: %v\n", testCase)
// }

//----------------------------------------------------------------------------------------------------
