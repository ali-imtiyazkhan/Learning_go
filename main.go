package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World")

	var confreneceName = "Go Conference"
	const confreneceTickets = 50
	var remainingTickets = confreneceTickets

	fmt.Println("welcome to", confreneceName, "booking application")
	fmt.Println("Number of tickets available :", remainingTickets)

	fmt.Println("Get your tickets now")

	// asking user for their name
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Your name")
	scanner.Scan()
	userName := scanner.Text()

	println("Hello", userName)

}
