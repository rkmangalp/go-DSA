package sorting2

import "sort"

// MergeSort recursively sorts an array using the merge sort algorithm.
func MergeSort(arr []int, low int, high int) {
	if low < high {
		mid := low + (high-low)/2 // Correct calculation of the midpoint
		// fmt.Printf("Msort called with low: %d, mid: %d, high: %d\n", low, mid, high)

		MergeSort(arr, low, mid)    // Recursively sort the first half
		MergeSort(arr, mid+1, high) // Recursively sort the second half
		Merge(arr, low, mid, high)  // Merge the two sorted halves
	}
}

// Merge combines two sorted subarrays into one sorted array.
func Merge(arr []int, low int, mid int, high int) {
	// fmt.Printf("Merge called with low: %d, mid: %d, high: %d\n", low, mid, high)

	temp := make([]int, 0, high-low+1) // Temporary slice to hold the merged elements
	left := low                        // Starting index for the left subarray
	right := mid + 1                   // Starting index for the right subarray

	// Merge the two subarrays into temp
	for left <= mid && right <= high {
		if arr[left] <= arr[right] { // Compare elements from both subarrays
			temp = append(temp, arr[left]) // Add the smaller element to temp
			left++                         // Move to the next element in the left subarray
		} else {
			temp = append(temp, arr[right]) // Add the smaller element to temp
			right++                         // Move to the next element in the right subarray
		}
	}

	// Append the remaining elements of the left subarray, if any
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	// Append the remaining elements of the right subarray, if any
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	// Copy the sorted elements back into the original array
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
	// fmt.Printf("Merged array: %v\n", arr[low:high+1])

}

// TEST CASE
// arr := []int{34, 7, 23, 32, 5, 62}
// 	fmt.Println("Given array is", arr)
// 	sorting2.MergeSort(arr, 0, len(arr)-1)
// 	fmt.Println("Sorted array is", arr)
//----------------------------------------------------------------------------------------------------

// 2. Merge Sort on Array of Strings

// Problem: Implement Merge Sort to sort an array of strings based on their alphabetical order.
// Example Input: ["apple", "orange", "banana", "pear"]
// Expected Output: ["apple", "banana", "orange", "pear"]

func MergeSortString(arr []string, low int, high int) {
	if low < high {
		mid := low + (high-low)/2
		MergeSortString(arr, low, mid)
		MergeSortString(arr, mid+1, high)
		MergeString(arr, low, mid, high)
	}
}

func MergeString(arr []string, low int, mid int, high int) {
	temp := make([]string, 0, high-low+1)
	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

// TEST CASE
// arr := []string{"apple", "orange", "banana", "pear"}
// 	fmt.Println("Given array is", arr)
// 	sorting2.MergeSortString(arr, 0, len(arr)-1)
// 	fmt.Println("Sorted array is", arr)
//----------------------------------------------------------------------------------------------------

// 3. Sort an Array of Objects using Merge Sort

// Problem: Implement Merge Sort to sort an array of objects based on a specific property.
// For instance, sort an array of Person objects by age.
// Example Input: [{"name": "Alice", "age": 25}, {"name": "Bob", "age": 20}, {"name": "Charlie", "age": 30}]
// Expected Output: [{"name": "Bob", "age": 20}, {"name": "Alice", "age": 25}, {"name": "Charlie", "age": 30}]

type Person struct {
	Name string
	Age  int
}

func MergeSortPerson(person []Person, low int, high int) {
	if low < high {
		mid := low + (high-low)/2
		MergeSortPerson(person, low, mid)
		MergeSortPerson(person, mid+1, high)
		MergePerson(person, low, mid, high)
	}
}

func MergePerson(person []Person, low int, mid int, high int) {
	temp := make([]Person, 0, high-low+1)
	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if person[left].Age <= person[right].Age {
			temp = append(temp, person[left])
			left++
		} else {
			temp = append(temp, person[right])
			right++
		}
	}
	for left <= mid {
		temp = append(temp, person[left])
		left++
	}
	for right <= high {
		temp = append(temp, person[right])
		right++
	}
	for i := low; i <= high; i++ {
		person[i] = temp[i-low]
	}
}

// TEST CASE
// persons := []sorting2.Person{
// 	{Name: "George", Age: 28},
// 	{Name: "Hannah", Age: 22},
// 	{Name: "Ivy", Age: 35},
// 	{Name: "Jack", Age: 30},
// }
// fmt.Println("Given array is:", persons)
// sorting2.MergeSortPerson(persons, 0, len(persons)-1)
// fmt.Println("Sorted array is:", persons)

//----------------------------------------------------------------------------------------------------

//4. Merge K Sorted Arrays

// Problem: Given K sorted arrays, merge them into a single sorted array.
// Example Input: [[1, 4, 7], [2, 5, 8], [3, 6, 9]]
// Expected Output: [1, 2, 3, 4, 5, 6, 7, 8, 9]

func MergeKsortedArray(arr [][]int) []int {
	var mergedArr []int

	for _, array := range arr {
		mergedArr = append(mergedArr, array...)
	}

	sort.Ints(mergedArr)
	return mergedArr

}
