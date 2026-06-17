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

	// continue: skip even numbers
	fmt.Println("Odd numbers:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// break with label (break outer loop)
	fmt.Println("Labeled break:")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break outer
			}
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}

	// continue with label
	fmt.Println("Labeled continue:")
outer2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue outer2
			}
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}
}