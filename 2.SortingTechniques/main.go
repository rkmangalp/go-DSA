package main

import (
	"fmt"
	"go-DSA/2.SortingTechniques/sorting2"
)

func main() {
	//copy paste the test cases to test the functions
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Initial arr: ", arr)
	sortedArr := sorting2.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("sorted arr: ", sortedArr) // Output: [11, 12, 22, 25, 64]

}
