package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 2, 1, 5, 3, 4}

	result := buildArray(arr)
	fmt.Println(result)
}

func buildArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] += 1000 * (nums[nums[i]] % 1000)
	}

	for i := 0; i < n; i++ {
		nums[i] /= 1000
	}

	return nums
}
