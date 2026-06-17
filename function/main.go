package main

import "fmt"

func main() {

	sayHello()

	var a int
	fmt.Println("Enter the value of a ")

	fmt.Scan(&a)

	var b int
	fmt.Println("Enter the value of b ")

	fmt.Scan(&b)

	var Sum = getSum(a,b)

	fmt.Println("The sum of ",a,"and",b,"is ",Sum)

	age := getTheUserAge()
	fmt.Println("Your age is : ", age)
	if age >= 18 {
		println("You are eligible for voting")
	} else {
		println("You are not eligible for voting")
	}

}

func sayHello() {
	println("Hello World")
	println("Welcome to Go Programming Language")
}

func getTheUserAge() int {
	var age int
	println("Enter your age here : ")
	fmt.Scan(&age)
	return age
}

func getSum(a int,b int) int {
	return a + b
}
