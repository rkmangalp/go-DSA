package hashingbasics

import "fmt"

func FrequencyArrEle() {

	// Input array of numbers
	numbers := []int{5, 2, 3, 4, 4, 2, 1}

	// Initialize a map to store the frequency of each number
	frequency := make(map[int]int)

	// Iterate through the array and count the occurrences of each number
	for _, num := range numbers {
		frequency[num]++
	}

	// Print the frequency of each number
	for num, count := range frequency {
		fmt.Printf("number %d appears %d times\n", num, count)
	}
}
