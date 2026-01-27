package main

import (
	"fmt"
)

// you can test any functions present in arrayProblems folder by calling it in main,go

func main() {

	arr := []int{1, 5, 2, 7, 18, 9, 10}
	result := largestElement(arr)
	fmt.Println(result)
}

func largestElement(nums []int) int {
	//your code goes here
	largest := nums[0] //3
	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return largest
}
