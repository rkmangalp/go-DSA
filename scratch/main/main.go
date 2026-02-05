package main

import "fmt"

func main() {
	students := map[string]int{"ravi": 90, "kiran": 100}
	students["mang"] = 95

	fmt.Println(students)
	delete(students, "kiran")

	fmt.Println("after delete:", students)

	//iterate

	for val := range students {
		fmt.Println(val)
	}
}
