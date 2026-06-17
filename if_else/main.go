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

}
