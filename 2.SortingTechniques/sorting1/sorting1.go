package sorting1

func SelectionSort(arr []int) []int { // Define a function that takes a slice of integers and returns a sorted slice

	n := len(arr)                  // Get the length of the array
	for i := 0; i < (n - 1); i++ { // Outer loop iterates from the start to the second last element
		min := i                 // Assume the current position is the minimum
		for j := i; j < n; j++ { // Inner loop to find the minimum element in the remaining unsorted portion
			if arr[j] < arr[min] { // If the current element is less than the assumed minimum
				min = j // Update the minimum position
			}
		}
		if min != i { // If the minimum element is not already in the current position
			arr[i], arr[min] = arr[min], arr[i] // Swap the current element with the found minimum element
		}
	}
	return arr // Return the sorted array
}

func BubbleSort(arr []int) []int {
	n := len(arr)                // Get the length of the array
	for i := n - 1; i > 0; i-- { // Outer loop: Iterate from the last element index to the second element index
		for j := 0; j < i; j++ { // Inner loop: Iterate from the first element index to the current `i`
			if arr[j] > arr[j+1] { // Compare current element with the next element
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap elements if they are out of order
			}
		}
	}
	return arr // Return the sorted array
}

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
