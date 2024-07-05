package main

import (
	"fmt"
	"go-DSA/3.arrayProblems/easy"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Original array is: %v\n", nums)
	roatatedArray := easy.RotatebyDnum1(nums, len(nums), 3)
	fmt.Printf("updated array is: %v", roatatedArray)

}
