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

	// for each loop
	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Println("the number is :", v)
	}

	// infinite loop
	// for {
	// 	fmt.Println("the number is :", i)
	// }

	// break and continue
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println("the number is :", i)
	}
}