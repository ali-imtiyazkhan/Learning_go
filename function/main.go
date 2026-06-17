package main

import "fmt"

func main() {

	sayHello()

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
