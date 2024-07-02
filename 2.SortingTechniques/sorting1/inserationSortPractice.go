package sorting1

import "fmt"

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ { // Outer loop: Starts from the second element and goes to the last element
		j := i
		for j > 0 && arr[j-1] > arr[j] { // Inner loop: Checks and swaps elements until the correct position is found
			arr[j-1], arr[j] = arr[j], arr[j-1] // Swap elements if they are out of order
			j--                                 // Move j backwards to continue comparing with previous elements
		}
	}
	return arr // Return the sorted array
}

// 1. Sort an Array of Integers
// Problem Statement:
// Implement the Insertion Sort algorithm to sort an array of integers in ascending order.
// Insertion Sort is a simple sorting algorithm that builds the final sorted array one item at a time.

func InsertionSort1(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		fmt.Printf("Iteration %d, key: %d, Array: %v\n", i, key, arr)
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		fmt.Printf("After iteration , key: %d, Array: %v\n", key, arr)
	}
	return arr
}

// arr := []int{12, 11, 13, 5, 6}
// fmt.Println("Initial Array:", arr)
// sortedArr := sorting1.InsertionSort1(arr)
// fmt.Println("Sorted Array:", sortedArr)
//----------------------------------------------------------------------------------------------------------------------------------------------------------

// 2. Sort an Array of Strings
// Problem Statement:
// Write a function that uses Insertion Sort to sort a slice of strings in alphabetical order.

func InsertionSortStrings(arr []string) []string {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// TEST CASE
// arr1 := []string{"banana", "apple", "cherry", "date"}
// 	sortedArr1 := sorting1.InsertionSortStrings(arr1)
// 	fmt.Println(sortedArr1)
//----------------------------------------------------------------------------------------------------------------------------------------------------------

// 3. Sort a Slice of Structs by a Field
// Problem Statement:
// Implement Insertion Sort to sort a slice of Student structs by the Grade field in ascending order.+

func InsertionSortStudents(students []Student) []Student {
	n := len(students)
	for i := 1; i < n; i++ {
		key := students[i]
		j := i - 1
		for j >= 0 && students[j].Grade > key.Grade {
			students[j+1] = students[j]
			j--
		}
		students[j+1] = key
	}
	return students
}

// TEST CASE
// students := []sorting1.Student{
// 	{Name: "Alice", Grade: 85},
// 	{Name: "Bob", Grade: 95},
// 	{Name: "Charlie", Grade: 75},
// 	{Name: "Dave", Grade: 90},
// }
// sortedStudents := sorting1.InsertionSortStudents(students)
// fmt.Println(sortedStudents)
//----------------------------------------------------------------------------------------------------------------------------------------------------------

// 4. Find the Kth Smallest Element After Sorting
// Problem Statement:
// Using Insertion Sort, sort an array of integers and then find the Kth smallest element.

func KthSmallestElementInsertion(arr []int, k int) int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr[k-1]
}
