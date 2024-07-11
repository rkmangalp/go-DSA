package easy

import (
	"fmt"
	"math"
	"sort"
)

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

//6. Move all the zeros to the end of the array

//Brute force

func MoveZeros(arr []int) []int {
	n := len(arr)
	temp := []int{}

	for _, num := range arr {
		if num != 0 {
			temp = append(temp, num)
		}
	}
	m := len(temp)
	for i := 0; i < m; i++ {
		arr[i] = temp[i]
	}
	for i := m; i < n; i++ {
		arr[i] = 0
	}
	return arr

}

// TEST CASE
// nums := []int{1, 0, 2, 3, 0, 4, 0, 5, 0, 6, 7}
// 	fmt.Printf("Original array is: %v\n", nums)
// 	endZeroArr := easy.MoveZeros(nums)
// 	fmt.Printf("updated array is: %v", endZeroArr)

//---------------------------------------------------------------------------------------------------------------------

//7. Linear Search

//Brute force

func LinearSearch(arr []int, num int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		if arr[i] == num {
			return i
		}
	}

	return -1
}

// TEST CASE
// arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
// 	num := 5
// 	index := easy.LinearSearch(arr, num)
// 	fmt.Printf("%d is present in %d index", num, index)

//---------------------------------------------------------------------------------------------------------------------

//8. Union of two sorted arrays

// Brute force
func UnionArray(arr1, arr2 []int) []int {

	// Initialize a map to store the frequency of elements
	freq := make(map[int]int)

	for _, value := range arr1 {
		freq[value]++
	}
	for _, value := range arr2 {
		freq[value]++
	}

	var Union []int

	for key := range freq {
		Union = append(Union, key)

	}
	sort.Ints(Union)
	return Union
}

// Optimal approach

func UnionArrayOpt(arr1, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)

	union := []int{}

	i := 0
	j := 0

	for i < n && j < m {
		if arr1[i] <= arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else {
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		}

	}
	for i < n {
		if len(union) == 0 || union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}

	for j < m {
		if len(union) == 0 || union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++

	}
	return union
}

//TEST CASE
// arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	arr2 := []int{2, 3, 4, 4, 5, 11, 12}
// 	Union := easy.UnionArrayOpt(arr1, arr2)

// 	fmt.Printf("Union of arr1 and arr2 is: %v", Union)

//---------------------------------------------------------------------------------------------------------------------

//9. Intersection of two sorted arrays (present in both)

//Brute force

func IntersecOfTwoArr(arr1, arr2 []int) []int {

	n1 := len(arr1)
	n2 := len(arr2)
	ansArr := []int{}
	visited := make(map[int]int)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if arr1[i] == arr2[j] && visited[j] == 0 {
				ansArr = append(ansArr, arr1[i])
				visited[j] = 1
				break
			}
			if arr1[i] < arr2[j] {
				break
			}
		}

	}
	return ansArr
}

//optimal Solution

func IntersecOfTwoArrOPT(arr1, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)
	i, j := 0, 0
	ansArr := []int{}

	for i < n && j < m {
		if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr1[i] {
			j++
		} else {
			ansArr = append(ansArr, arr1[i])
			i++
			j++
		}
	}
	return ansArr
}

//TEST CASE
// arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	arr2 := []int{2, 3, 4, 4, 5, 11, 12}
// 	result := easy.IntersecOfTwoArr(arr1, arr2)

// 	fmt.Printf("Union of arr1 and arr2 is: %v\n", result)

//---------------------------------------------------------------------------------------------------------------------
// 10 . Find the missing number in an array

// Problem Statement: Given an integer N and an array of size N-1 containing N-1 numbers between 1 to N.
// Find the number(between 1 to N), that is not present in the given array.

//Bruteforce method

func FindMissingNumBrute(arr []int, n int) int {

	for i := 1; i <= n; i++ {
		found := false
		for j := 0; j < n-1; j++ {
			if arr[j] == i {
				found = true
				break
			}

		}
		if !found {
			return i
		}

	}
	return -1
}

//Better Method

func FindMissingNumBetter(arr []int, n int) int {
	hash := make(map[int]int)

	for _, num := range arr {
		hash[num]++
	}
	for i := 1; i <= n; i++ {
		if hash[i] == 0 {
			return i
		}
	}
	return -1
}

//Optimal approach  (using summation )

func FindMissingNumOpt1(arr []int, n int) int {
	sum := (n * (n + 1)) / 2
	arrSum := 0

	for _, num := range arr {
		arrSum = arrSum + num
	}
	return sum - arrSum

}

//Optimal approach  (using XOR method )

func FindMissingNumOpt2(arr []int, n int) int {
	xor1 := 0
	xor2 := 0

	for i := 1; i <= n; i++ {
		xor1 ^= i
	}

	for _, num := range arr {
		xor2 ^= num
	}

	return xor1 ^ xor2
}

//Test case
// arr := []int{1, 2, 3, 4, 5, 6, 8, 9}
// n := 9
// result := easy.FindMissingNumOpt2(arr, n)

// fmt.Printf("The missing number is %d in array %v\n", result, arr)
//---------------------------------------------------------------------------------------------------------------------

// 11. Count Maximum Consecutive One's in the array

// Problem Statement: Given an array that contains only 1 and 0 return the count of
// maximum consecutive ones in the array.

func MaxConsecutiveOnes(arr []int) int {
	max := 0
	count := 0

	for _, num := range arr {

		if num == 1 {
			count++
			max = int(math.Max(float64(max), float64(count)))
		} else {
			count = 0
		}
	}
	return max
}

//Test case
// arr := []int{0, 1, 0, 1, 1, 1, 1}
// result := easy.MaxConsecutiveOnes(arr)

// fmt.Printf("The max consecutive ones is %d in array %v\n", result, arr)

//---------------------------------------------------------------------------------------------------------------------

//12. Find the number that appears once, and the other numbers twice

//Problem Statement: Given a non-empty array of integers arr, every element appears twice except for one.
//Find that single one.

// Brute force- using linear search and keeping a count of each element

//Brute approach

func FindNumAppearOnesBrute(arr []int) int {
	n := len(arr)

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] {
				count++
			}

		}
		if count == 1 {
			return arr[i]
		}
	}

	return -1

}

// Better approach (using hash maps)

func FindNumAppearOnesBetter(arr []int) int {

	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	for num, count := range counts {
		if count == 1 {
			return num
		}
	}
	return -1
}

func FindNumAppearOnesOptimal(arr []int) int {

	xor := 0
	for i := 0; i < len(arr); i++ {
		xor = xor ^ arr[i]
	}
	return xor
}

//Test case
// arr := []int{1, 1, 2, 3, 3, 4, 4}
// result := easy.FindNumAppearOnesOptimal(arr)

// fmt.Printf("The number %d appears ones in the array %v\n", result, arr)

//---------------------------------------------------------------------------------------------------------------------

//13. Longest Subarray with given Sum K(Positives)
//Problem Statement: Given an array and a sum k, we need to print the length of the longest subarray that sums to k.

// Brute force
func LongestSubArrayBrute(arr []int, k int) (int, []int) {
	n := len(arr)
	maxLen := 0
	start := -1
	end := -1
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == k {
				length := j - i + 1
				if length > maxLen {
					maxLen = length
					start = i
					end = j
				}
			}
		}
	}
	return maxLen, arr[start : end+1]
}

// Better Approach
func LongestSubArrayBetter(arr []int, k int) int {
	n := len(arr) // size of the array

	// preSumMap will store the prefix sums and their indice
	preSum := make(map[int]int)
	sum := 0    // sum of elements seen so far
	maxLen := 0 // maximum length of subarray with sum equal to k

	// iterate through the array
	for i := 0; i < n; i++ {
		sum += arr[i] // add the current element to the running sum

		if sum == k {
			maxLen = max(maxLen, i+1)
		}
		// calculate the remaining sum that needs to be subtracted to get k
		rem := sum - k

		// if rem exists in the map, it means there is a subarray with sum equal to k
		if index, found := preSum[rem]; found {
			length := i - index
			maxLen = max(length, maxLen)
		}
		// if sum is not already in the map, add it with its index
		if _, found := preSum[sum]; !found {
			preSum[sum] = i
		}

	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LongestSubArrayOptimal(arr []int, k int) (int, []int) {
	n := len(arr)
	left, right := 0, 0
	sum := arr[0]
	maxLen := 0
	start := 0
	end := 0

	for right < n {
		for left <= right && sum > k {
			sum -= arr[left]
			left++
		}
		if sum == k {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
				start, end = left, right
			}

		}

		right++
		if right < n {
			sum += arr[right]
		}

	}

	return maxLen, arr[start : end+1]
}

//Test case
// a := []int{2, 3, 5, 1, 9}
// k := 10
// len, arr := easy.LongestSubArrayOptimal(a, k)
// fmt.Printf("The length of the longest subarray is: %d and the subarray is %v", len, arr)
//---------------------------------------------------------------------------------------------------------------------
