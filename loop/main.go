package main

import "fmt"

func main() {
	// for loop
	for i := 1; i < 10; i++ {
		fmt.Println("the number is :", i)
	}

	// while loop equivalent
	i := 1
	for i < 10 {
		fmt.Println("the number is :", i)
		i++
	}
}