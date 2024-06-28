package patterns

import "fmt"

//Pattern 1

// *****
// *****
// *****
// *****
// *****

func Pattern1(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

//Pattern 2

// *
// **
// ***
// ****
// *****

func Pattern2(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

//Pattern 3

// 1
// 12
// 123
// 1234
// 12345

func Pattern3(num int) {

	for i := 1; i < num+1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)

		}
		fmt.Println()
	}

}

//Pattern 4

// 1
// 22
// 333
// 4444
// 55555

func Pattern4(num int) {

	for i := 1; i < num+1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)

		}
		fmt.Println()
	}

}

//Pattern 5

// ******
// *****
// ****
// ***
// **
// *

func Pattern5(num int) {

	for i := num; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print("*")

		}
		fmt.Println()
	}

}

//Pattern 6
// 12345
// 1234
// 123
// 12
// 1

func Pattern6(num int) {

	for i := num; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)

		}
		fmt.Println()
	}

}

//Pattern 7

//      *
//     ***
//    *****
//   *******
//  *********
func Pattern7(num int) {

	// iterate for number of rows
	for i := 0; i < num; i++ {
		space := num - i - 1
		star := 2*i + 1

		//print the space
		for j := 0; j <= space; j++ {
			fmt.Print(" ")
		}

		//print the star
		for j := 0; j < star; j++ {
			fmt.Print("*")
		}
		//print the space
		for j := 0; j <= space; j++ {
			fmt.Print(" ")
		}
		fmt.Println()

	}

}

//Pattern 8

//  *********
//   *******
//    *****
//     ***
//      *
func Pattern8(num int) {

	// iterate for number of rows
	for i := 0; i < num; i++ {
		star := 2*num - (2*i + 1)

		//print the space
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}

		//print the star
		for j := 0; j < star; j++ {
			fmt.Print("*")
		}
		//print the space
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()

	}
}

//Pattern 9
func Pattern9(num int) {
	// combine pattern 7 & pattern 8
}

//Pattern 10

// *
// **
// ***
// ****
// *****
// ****
// ***
// **
// *
func Pattern10(num int) {

	// iterate for rows 2*n -1
	for i := 0; i <= 2*num-1; i++ {
		stars := i
		if i > num {
			stars = 2*num - i
		}
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

}

//Pattern 11

// 1
// 01
// 101
// 0101
// 10101

func Pattern11(num int) {

	start := 1
	for i := 0; i < num; i++ {

		if i%2 == 0 {
			start = 1
		} else {
			start = 0
		}
		for j := 0; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()
	}

}

//Pattern 12

// Need to fix this gap ******

// 1         1
// 12       21
// 123     321
// 1234   4321
// 12345 54321
func Pattern12(num int) {
	space := 2 * (num - 1)
	for i := 1; i <= num; i++ {

		//Prints numbers
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		// print space
		for j := 0; j <= space; j++ {
			fmt.Print(" ")
		}
		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
		space = space - 2

	}

}

//Pattern 13

// 1
// 2 3
// 4 5 6
// 7 8 9 10
// 11 12 13 14 15

func Pattern13(num int) {
	n := 1
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(n, " ")

			n++
		}
		fmt.Println()
	}

}

//Pattern 14

// A
// A B
// A B C
// A B C D
// A B C D E

func Pattern14(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", 'A'+rune(j))
		}
		fmt.Println()
	}
}

//Pattern 15

// A B C D E F
// A B C D E
// A B C D
// A B C
// A B
// A

func Pattern15(n int) {
	for i := n; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", 'A'+rune(j))
		}
		fmt.Println()
	}
}

//Pattern 16
// A
// B B
// C C C
// D D D D
// E E E E E
func Pattern16(n int) {
	ch := 'A'
	for i := 0; i < n; i++ {

		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", ch+rune(i))
		}
		fmt.Println()
	}
}

//Pattern 17

// Need to fix this code

func Pattern17(n int) {
	for i := 0; i < n; i++ {

		// space
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		//initialize char
		ch := 'A'

		mid := (2 * i) / 2

		for j := 0; j <= 2*i; j++ {
			fmt.Printf("%c ", ch)

			if j < mid {
				ch++
			} else {
				ch--
			}

		}
		//space
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// Pattern 18
// E
// D E
// C D E
// B C D E
// A B C D E
func Pattern18(n int) {

	for i := 0; i < n; i++ {

		char := 'A' + rune(n-1)
		// Start character for each line, which decreases as i increases
		startChar := char - rune(i)
		for j := 0; j <= i; j++ {
			// Print the character, incremented by j
			fmt.Printf("%c ", startChar+rune(j))
		}
		fmt.Println()
	}
}

// pattern 19

// **********
// ****  ****
// ***    ***
// **      **
// *        *
// *        *
// **      **
// ***    ***
// ****  ****
// **********

func Pattern19(n int) {

	// loop the first part of the pattern

	for i := 0; i < n; i++ {

		// stars
		for j := n; j > i; j-- {
			fmt.Print("*")
		}

		//space
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}

		//stars
		for j := n; j > i; j-- {
			fmt.Print("*")
		}

		fmt.Println()
	}

	// loop for 2nd part of the pattern

	for i := 0; i < n; i++ {

		// stars
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		//space
		for j := 0; j < 2*(n-i)-2; j++ {
			fmt.Print(" ")
		}

		//stars
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

}

// Pattern 20

// *        *
// **      **
// ***    ***
// ****  ****
// **********
// ****  ****
// ***    ***
// **      **
// *        *

func Pattern20(n int) {

	// first half

	for i := 0; i < n; i++ {

		// stars
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		//space
		for j := 0; j < 2*(n-i)-2; j++ {
			fmt.Print(" ")
		}

		//stars
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	// second half

	for i := 1; i < n; i++ {

		// stars
		for j := n; j > i; j-- {
			fmt.Print("*")
		}

		//space
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}

		//stars
		for j := n; j > i; j-- {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

//Pattern 21

// ****
// *  *
// *  *
// ****

func Pattern21(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}

// Pattern 22

func Pattern22(n int) {
	end := 2*n - 1
	for i := 0; i < end; i++ {
		for j := 0; j < end; j++ {
			top, left, right, bottom := i, j, (2*n-2)-j, (2*n-2)-i

			fmt.Print(n - min(min(top, bottom), min(left, right)))
		}
		fmt.Println()
	}
}
