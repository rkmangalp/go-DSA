package main

import (
	"fmt"
	"go-DSA/2.SortingTechniques/sorting1"
)

func main() {

	// testcase for Selection sort And Bubble sort
	arr := []int{13, 46, 24, 52, 20, 9}
	sortedArr := sorting1.InsertionSort(arr)
	fmt.Println(sortedArr)
}
