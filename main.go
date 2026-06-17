package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	var confreneceName = "Go Conference"
	const confreneceTickets = 50
	var remainingTickets = confreneceTickets

	fmt.Println("welcome to", confreneceName, "booking application")
	fmt.Println("Number of tickets available :", remainingTickets)

	fmt.Println("Get your tickets now")

	var userName string
	// asking user for their name

	fmt.Println("Enter Your name")
	fmt.Scan(&userName)

	println("Hello", userName)

}
