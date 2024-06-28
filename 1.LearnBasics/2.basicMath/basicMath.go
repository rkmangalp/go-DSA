package basicMath

import (
	"fmt"
	"math"
	"sort"
)

// Euclidean algorithm to find the GCD
func EuclideanGCD(a int, b int) int {

	for (a > 0) && (b > 0) {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	if a == 0 {
		return b
	}
	return a
}

// Greatest common divisor
func Gcd(n1 int, n2 int) int {

	gcd := 0

	if max(n1, n2)%min(n1, n2) == 0 {
		gcd = min(n1, n2)
	}
	//loop till min number
	for i := 1; i <= min(n1, n2); i++ {
		if n1%i == 0 && n2%i == 0 {
			gcd = i
		}

	}
	return gcd
}

// check for prime
func PrimeNum(n int) {
	prime := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			prime = append(prime, i)
		}

	}
	if len(prime) == 2 {
		fmt.Printf("%d is a prime number because it has two divisors %v", n, prime)
	} else {
		fmt.Printf("%d is a composite number because it has %d two divisors %v", n, len(prime), prime)

	}
}

// Divisors optimal solution to reduce time complexity by iterating till aquare root of n
func DivisorsOptimal(n int) {
	divisors := []int{}
	end := int(math.Sqrt(float64(n)))
	for i := 1; i <= end; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if n/i != i {
				num := n / i
				divisors = append(divisors, num)
			}
		}
	}
	sort.Ints(divisors)
	fmt.Printf("The divisors of %d are %v", n, divisors)
}

// Print all divisors of Given number
func Divisors(n int) {
	divisors := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}

	}
	fmt.Printf("The divisors of %d are %v", n, divisors)
}

// If the summation of cubes of digits in the number is equal to number then its called armstrong number
func ArmstrongNum(n int) {
	// 35
	sum := 0
	number := n
	for n > 0 {
		num := n % 10
		sum = sum + int(math.Pow(float64(num), 3))
		n = n / 10
	}
	if number == sum {
		fmt.Printf(" The sum of cube of digits of number %d is %d and therefor it is an Armstrong number", number, sum)

	} else {
		fmt.Printf(" The sum of cube of digits of number %d is %d and therefor it is not an Armstrong number", number, sum)
	}
}

// checks if the num is palindrome
func Palindrome(n int) {

	revNum := 0
	num := n
	for n > 0 {
		num := n % 10
		revNum = (revNum * 10) + num
		n = n / 10

	}
	if revNum == num {
		fmt.Printf(" The reverse of %d is %d and therefor it is a palindrome", num, revNum)
	} else {
		fmt.Printf(" The reverse of %d is %d and therefor it is not a palindrome", num, revNum)
	}

}

// returns the reverse of the given digit
func ReverseDigits(n int) int {

	revNum := 0

	for n > 0 {
		num := n % 10
		revNum = (revNum * 10) + num
		n = n / 10
	}
	return revNum
}

// counts the number of digits
func CountDigits(n int) int {

	// brute force approach
	if n == 0 {
		return 1
	}

	count := 0

	for n > 0 {
		count++
		n = n / 10
	}
	return count

	// optimal approach
	// return int(math.Floor(math.Log10(float64(n))) + 1)
}
