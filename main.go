package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Hello World")
	fmt.Printf("welcome to %v booking application.\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}
