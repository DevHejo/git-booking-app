package main

import "fmt"

func main() {
	// declaring variables
	// var conferenceName = "Go Conference"
	// const conferenceTickets = 50
	// var remainingTickets = 50

	// another way to decalre variables
	// var conferenceName string = "Go Conference"
	// const conferenceTickets int = 50
	// var remainingTickets uint = 50

	conferenceName := "Go Conference" // this is same as conferenceName = "Go Conference"
	const conferenceTickets int = 50  // the above method cannot be used while declaring constants
	var remainingTickets uint = 50    // the above method connot be used  while declaring types

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Hello World")
	fmt.Printf("welcome to %v booking application.\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int

	userName = "Jack"
	userTickets = 2
	fmt.Printf("Hello %v\n", userName)
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
