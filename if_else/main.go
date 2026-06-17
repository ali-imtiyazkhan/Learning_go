package main

import "fmt"

func main() {

	var age int
	fmt.Println("Enter your age here : ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("you can't vote")
	}

	// if-else if ladder
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// if with short statement
	if num := 15; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// nested if
	x := 10
	y := 20
	if x == 10 {
		if y == 20 {
			fmt.Println("x=10 and y=20")
		}
	}
}
