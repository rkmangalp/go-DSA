package sorting1

//Selection sort practice problems
//--------------------------------------------------------------------------------------------------'
// Problem 1: Sorting Structs

// Given a list of structs representing students, each with a name and grade,
// implement selection sort to sort the list by grades in descending order.
type Student struct {
	Name  string
	Grade int
}

func SelectionSortStudents(students []Student) []Student {
	n := len(students)
	for i := 0; i < n-1; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if students[max].Grade < students[j].Grade {
				max = j
			}
		}
		if max != i {
			students[max], students[i] = students[i], students[max]
		}
	}
	return students
}

// TEST CASE:- (paste below code in main.go)

// students := []sorting1.Student{
// 	{Name: "Alice", Grade: 85},
// 	{Name: "Bob", Grade: 95},
// 	{Name: "Charlie", Grade: 75},
// 	{Name: "Dave", Grade: 90},
// }
// sortedStudents := sorting1.SelectionSortStudents(students)
// fmt.Println(sortedStudents)
// Expected output: [{Bob 95} {Dave 90} {Alice 85} {Charlie 75}]
//--------------------------------------------------------------------------------------------------'

// Problem 2: Finding the k-th Smallest Element

// Using selection sort, find the k-th smallest element in an unsorted array of integers.
// Note that you don't need to fully sort the array to find the k-th smallest element.

func KthSmallestElement(arr []int, k int) int {
	n := len(arr)
	for i := 0; i < k; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}

		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}

	}
	return arr[k-1]
}

//TEST CASE

// arr := []int{7, 10, 4, 3, 20, 15}
// k := 3
// result := sorting1.KthSmallestElement(arr, k)
// fmt.Println("The", k, "rd smallest element is:", result)

//--------------------------------------------------------------------------------------------------'

// Problem 3: Sorting Strings

// Implement a function that sorts a slice of strings in alphabetical order using selection sort.

func SelectionSortStrings(arr []string) []string {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min][0] > arr[j][0] {
				min = j
			}
			if min != i {
				arr[min], arr[i] = arr[i], arr[min]
			}
		}

	}
	return arr
}

//TEST CASE

// input := []string{"banana", "apple", "cherry", "date"}
// output := sorting1.SelectionSortStrings(input)
// fmt.Println(output)
// Expected output: ["apple", "banana", "cherry", "date"]

//--------------------------------------------------------------------------------------------------

// Problem 4: Sorting an Array of Pairs

// Given an array of pairs where each pair consists of two integers,
// sort the array by the first integer of each pair using selection sort. If the first integers are equal, sort by the second integer.

type Pair struct {
	First  int
	Second int
}

func SelectionSortPairs(arr []Pair) []Pair {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min].First > arr[j].First {
				min = j
			}
			if arr[min].First == arr[j].First {
				if arr[min].Second > arr[j].Second {
					min = j
				}
			}
			if min != i {
				arr[min], arr[i] = arr[i], arr[min]
			}
		}

	}
	return arr
}

//TEST CASE

// input := []sorting1.Pair{{3, 4}, {1, 2}, {3, 1}, {2, 5}}
// output := sorting1.SelectionSortPairs(input)
// fmt.Println(output)
// Expected output: [{1, 2}, {2, 5}, {3, 1}, {3, 4}]
//--------------------------------------------------------------------------------------------------
