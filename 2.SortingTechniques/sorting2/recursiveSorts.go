package sorting2

import "fmt"

// RecursiveBubbleSort sorts an array using the bubble sort algorithm recursively.
// arr is the array to be sorted, and n is the size of the array (or the current unsorted portion).
func RecursiveBubbleSort(arr []int, n int) []int {
	// Base case: If the array has only one element, it is already sorted.
	if n == 1 {
		return arr
	}

	// Perform one pass of bubble sort to move the largest element in the current subarray to the end.
	for i := 0; i < n-1; i++ {
		// If the current element is greater than the next element, swap them.
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			fmt.Printf("iteration: %d arr: %v and n value is: %d\n", i, arr, n)
		}
	}

	// Recursively call RecursiveBubbleSort to sort the rest of the array.
	// The largest element is now at the end of the array, so we reduce the problem size by 1 (n-1).
	RecursiveBubbleSort(arr, n-1)

	return arr

}

// TEST CASE
// arr := []int{64, 34, 25, 12, 22, 11, 90}
// n := len(arr)
// fmt.Println("Given array is:", arr)
// sortedArr := RecursiveBubbleSort(arr, n)
// fmt.Println("Sorted array is:", sortedArr)
//----------------------------------------------------------------------------------------------------

// RecursiveInsertionSort sorts an array using the insertion sort algorithm recursively.
// arr is the array to be sorted, i is the current index, and n is the size of the array.
func RecursiveInsertionSort(arr []int, i int, n int) []int {
	// Base case: If we have sorted all elements, return the array.
	if i == n {
		return arr
	}

	// Move elements of arr[0..i-1], that are greater than arr[i], to one position ahead
	// of their current position.
	for i > 0 && arr[i-1] > arr[i] {
		arr[i], arr[i-1] = arr[i-1], arr[i]
		i--
	}

	// Recursively call the function to sort the next element.
	RecursiveInsertionSort(arr, i+1, n)

	return arr
}

// TEST CASE
// arr := []int{12, 11, 13, 5, 6} // Example input array
// n := len(arr) // Size of the array
// fmt.Println("Given array is:", arr)
// sortedArr := RecursiveInsertionSort(arr, 1, n) // Start from index 1
// fmt.Println("Sorted array is:", sortedArr) // Output the sorted array
//----------------------------------------------------------------------------------------------------
