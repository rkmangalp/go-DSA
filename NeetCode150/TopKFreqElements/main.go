package main

import (
	"fmt"
	"sort"
)

// Top K Frequent Elements
// Given an integer array nums and an integer k, return the k most
// frequent elements. You may return the answer in any order.
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]

// Example 3:
// Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2
// Output: [1,2]

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.

// Follow up: Your algorithm's time complexity must be better than
// O(n log n), where n is the array's size.
func main() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	fmt.Println(topkFrequencyBetter(nums, 2))
}

// BRUTE FORCE
// Count frequency of every element.
// Find the element with highest frequency.
// Remove it.
// Repeat this process k times.

func topKFrequentBrute(nums []int, k int) []int {
	freq := make(map[int]int)

	//build an map
	for _, val := range nums {
		freq[val]++
	}
	result := []int{}
	// [1:4,2:3,3:1]
	// find max element k times
	for i := 0; i < k; i++ {
		maxNum := 0
		maxFreq := 0

		for num, f := range freq {
			if f > maxFreq {
				maxFreq = f
				maxNum = num
			}
		}
		result = append(result, maxNum)
		delete(freq, maxNum)
	}
	return result
}

// Better
// Count frequency (same as brute)
// Convert map → slice
// Sort slice by frequency (descending)
// Pick first k

func topkFrequencyBetter(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	type pair struct {
		num  int
		freq int
	}

	arr := []pair{}

	for num, f := range freq {
		arr = append(arr, pair{num, f})
	}

	sort.Slice(arr, func(i int, j int) bool {
		return arr[i].freq > arr[j].freq
	})
	fmt.Println(arr)

	result := []int{}

	for i := 0; i < k; i++ {
		result = append(result, arr[i].num)
	}
	return result
}

// Optimal solution (bucket sort)
// Frequency of any element can range from 1 to n
// We don’t need full sorting
// We can group numbers by frequency
// 🔥 Bucket Sort (based on frequency)

func topKFrequentOptimal(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	bucket := make([][]int, len(nums)+1)

	for num, f := range freq {
		bucket[f] = append(bucket[f], num)
	}

	result := []int{}
	for i := len(bucket) - 1; i < 0 && len(result) < k; i-- {
		for _, num := range bucket[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}

	return result
}
