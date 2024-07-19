package main

import (
	"fmt"
	"go-DSA/3.arrayProblems/medium"
)

// you can test any functions present in arrayProblems folder by calling it in main,go

func main() {
	//Test case
	tests := [][]int{
		{3, 3, 4, 2, 4, 4, 2, 4, 4}, // Test case 1
		{1, 1, 1, 1, 1, 1, 1},       // Test case 2
		{2, 2, 1, 1, 2, 2, 2},       // Test case 3
		{1, 1, 1, 2, 2, 2, 2},       // Test case 4
		{5, 5, 5, 1, 2, 5, 3, 5},    // Test case 5
	}

	for i, arr := range tests {
		result := medium.MajorityElementInArrOptimal(arr)
		fmt.Printf("Test case %d: Input: %v, Majority Element: %d\n", i+1, arr, result)
	}

}
