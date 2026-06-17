package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value at pointer p:", *p)

	*p = 20
	fmt.Println("New value of x:", x)
}
