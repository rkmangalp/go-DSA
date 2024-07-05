package easy

import "fmt"

// 1. find the largest element in an array

// Brute force method
func LargestElementinArray(arr []int) int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[n-1]
}

// Optimal Method

func LargestElementinArray2(arr []int) int {
	largest := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	return largest
}

//Test case (run the below in main.go)
// arr := []int{3, 2, 1, 5, 2}
// 	res := easy.LargestElementinArray(arr)
// 	fmt.Println(res)

//----------------------------------------------------------------------------------------------------

// 2. find the second largest element in array

// Brute force method - This will work only if there are no duplicates in the array

func SecondLargestinArray1(arr []int) int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[n-2]
}

// Better Approach - to find the  second largest in array
// {1, 2, 4, 7, 7, 5}
func SecondLargestinArray2(arr []int) int {
	largest := arr[0]
	slargest := -1

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	fmt.Printf("Largest number is set to: %d\n", largest)
	for i := 0; i < len(arr); i++ {
		if arr[i] > slargest && arr[i] < largest {
			slargest = arr[i]
		}
	}
	return slargest
}

// optimal approcah - to find the second largest in array

func SecondLargestinArray3(arr []int) int {
	largest := arr[0]
	slargest := -1

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			slargest = largest
			largest = arr[i]
		} else if arr[i] < largest && arr[i] > slargest {
			slargest = arr[i]
		}

	}
	return slargest

}

//----------------------------------------------------------------------------------------------------

// 3. Remove duplicates in place from a sorted array and return the number of unique elements

//Brute force

func RemoveDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	uniqueCount := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[uniqueCount] = arr[i]
			uniqueCount++
		}

	}
	return uniqueCount
}

//Optimal approach - 2 pointer method

func RemoveDuplicates1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[i] {
			arr[i+1] = arr[j]
			i++
		}
	}
	return i + 1
}

// TEST CASE
// nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
// 	uniqueCount := easy.RemoveDuplicates(nums)
// 	fmt.Printf("Number of unique elements: %d\n", uniqueCount)
// 	fmt.Printf("Array after removing duplicates: %v\n", nums[:uniqueCount])

//----------------------------------------------------------------------------------------------------
// 4. Left rotate the array by one place

// brute force can be solved by declaring an empty array

func LeftRotateBrute(arr []int) []int {
	n := len(arr)
	if len(arr) == 0 {
		return arr
	}
	rotateArr := make([]int, n)
	for i := 0; i < n-1; i++ {
		rotateArr[i] = arr[i+1]
	}
	rotateArr[n-1] = arr[0]
	return rotateArr
}

//optimal approach

func LeftRotate(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	temp := arr[0]
	n := len(arr)
	for i := 1; i < n; i++ {
		arr[i-1] = arr[i]
	}
	arr[n-1] = temp
	return arr
}

//TEST CASE
// nums := []int{1, 2, 3, 4, 5}
// 	leftRotateArr := easy.LeftRotateBrute(nums)
// 	fmt.Printf("updated array is: %v", uniqueCount)

//----------------------------------------------------------------------------------------------------

// 5. Left rotate the array by D places

//Brute force method

func RotatebyDnum(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	d = d % n // this will avoid unnecessary rotations
	temp := append([]int{}, arr[:d]...)
	//or you can use "copy(temp, arr[:d])"
	fmt.Println(temp)
	for i := d; i < n; i++ {
		arr[i-d] = arr[i]
	}

	for i := n - d; i < n; i++ {
		arr[i] = temp[i-d-1]
	}
	return arr
}

// Time complexity : O(n+d)
// spcae complexity: O(d)

// optimal solution - instead of using a temp array (make space complexity to O(1))

func RotatebyDnum1(arr []int, n int, d int) []int {
	Reverse(arr, 0, d-1)
	Reverse(arr, d, n-1)
	Reverse(arr, 0, n-1)
	return arr
}
func Reverse(arr []int, start int, end int) []int {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--

	}
	return arr
}

// TEST CASE
// nums := []int{1, 2, 3, 4, 5, 6, 7}
// 	fmt.Printf("Original array is: %v\n", nums)
// 	ReveredArray := easy.RotatebyDnum(nums, 3)
// 	fmt.Printf("updated array is: %v", uniqueCount)

//----------------------------------------------------------------------------------------------------
