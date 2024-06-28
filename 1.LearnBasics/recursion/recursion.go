package recursion

import "fmt"

func Fibonacci(n int) (num int) {
	if n <= 1 {
		return n
	}
	last := Fibonacci(n - 1)
	slast := Fibonacci(n - 2)

	return last + slast
}

//check Palindrome
func PalindromeRecurssion(s string, start, end int) bool {
	if start >= end {
		return true
	}

	return PalindromeRecurssion(s, start+1, end-1)
}

//Reverse an Array using single pointer
func ReverseAnArrayRecurssion2(arr []int, i int) {
	if i >= len(arr)/2 {
		return
	}
	arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]

	ReverseAnArrayRecurssion2(arr, i+1)
}

//Reverse an Array
func ReverseAnArrayRecurssion(arr []int, start int, end int) {

	if start >= end {
		return
	}
	arr[start], arr[end] = arr[end], arr[start]
	ReverseAnArrayRecurssion(arr, start+1, end-1)
}

//Factorial of N
func FactorialRecursion(n int) (sum int) {
	if n == 1 {
		return 1
	}

	return n * FactorialRecursion(n-1)
}

//Print sum of name in recusion, Functional way
func SumNumRecursion2(n int) (sum int) {
	if n == 0 {
		return 0
	}

	return n + SumNumRecursion2(n-1)

}

//Print sum of name in recusion, Parameter wise
func SumNumRecursion(n, sum int) {
	if n < 0 {
		fmt.Println(sum)
		return
	}

	SumNumRecursion(n-1, sum+n)
}

//Print reverse N times
func PrintReverseNum(i, n int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	PrintReverseNum(i-1, n)
}

//Print n times using recursion
func PrintNtimes(i int, n int) {
	if i > n {
		return
	}
	fmt.Println(n)

	PrintNtimes(i+1, n)
}

func PritnRecursion() {
	count := 0 // need to remove this and initialise gobally
	if count == 3 {
		return
	}
	fmt.Println(1)
	count++
	PritnRecursion()

}
