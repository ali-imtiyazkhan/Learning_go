package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var email string
	var age int

	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	if name == "" || len(name) < 2 ||
		email == "" || !strings.Contains(email, "@") ||
		age < 18 {

		fmt.Println("Please correctly fill all the fields")
	} else {
		fmt.Println("Name :", name)
		fmt.Println("Email:", email)
		fmt.Println("Age  :", age)
	}
}