package main

import (
	"Go_learning/healper"
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World")

	healper.Helper()

	confreneceName := "Go Conference"
	const confreneceTickets = 50
	remainingTickets := confreneceTickets

	fmt.Println("welcome to", confreneceName, "booking application")
	fmt.Println("Number of tickets available :", remainingTickets)

	fmt.Printf("confrenceName is type %T\n", confreneceName)
	fmt.Printf("conreneceTickets is type %T\n", confreneceTickets)
	fmt.Printf("remainingTickets is type %T\n", remainingTickets)

	fmt.Println("Get your tickets now")

	// asking user for their name
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Your name")
	scanner.Scan()
	userName := scanner.Text()

	println("Hello", userName)

}
