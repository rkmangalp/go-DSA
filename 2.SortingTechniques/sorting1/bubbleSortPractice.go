package sorting1

import "fmt"

// Bubble sort Practice problems
//-------------------------------------------------------------------------------------------------------------------------------------
// 1. Basic Bubble Sort
// Problem:
// Implement the bubble sort algorithm to sort an array of integers in ascending order.

func BubbleSort1(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// TEST CASE-
// arr := []int{64, 25, 12, 22, 11}
// 	sortedArr := sorting1.BubbleSort1(arr)
// 	fmt.Println(sortedArr) // Output: [11, 12, 22, 25, 64]

//-------------------------------------------------------------------------------------------------------------------------------------
// 2. Bubble Sort with Optimizations
// Problem:
// Modify the bubble sort algorithm to optimize it by stopping early if the array is already sorted.

func OptimizedBubbleSort(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {

			if arr[j] > arr[j+1] {
				fmt.Println(j)
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

// Test Case
//   arr1 := []int{64, 25, 12, 22, 11}
//   sortedArr1 := OptimizedBubbleSort(arr1)
//   fmt.Println(sortedArr1) // Expected: [11, 12, 22, 25, 64]

//-------------------------------------------------------------------------------------------------------------------------------------

// 3. Bubble Sort for Strings
func BubbleSortStrings(arr []string) []string {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// TEST CASE
// arr1 := []string{"banana", "apple", "cherry", "date"}
// sortedArr1 := BubbleSortStrings(arr1)
// fmt.Println(sortedArr1)
//-------------------------------------------------------------------------------------------------------------------------------------

// 4. Bubble Sort for Structs

type Students struct {
	Name  string
	Grade int
}

func BubbleSortStudents(students []Student) []Student {
	n := len(students)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if students[j].Grade > students[j+1].Grade {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
	return students
}

//TEST CASE-
// students := []sorting1.Student{
// 	{Name: "Alice", Grade: 85},
// 	{Name: "Bob", Grade: 95},
// 	{Name: "Charlie", Grade: 75},
// 	{Name: "Dave", Grade: 90},
// }
// sortedStudents := sorting1.BubbleSortStudents(students)
// fmt.Println(sortedStudents)
