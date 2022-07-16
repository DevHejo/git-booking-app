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

	conferenceName := "Go Conference" // this is same as var conferenceName = "Go Conference"
	const conferenceTickets uint = 50 // the above method cannot be used while declaring constants
	var remainingTickets uint = 50    // the above method connot be used  while declaring types

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Hello World")
	fmt.Printf("welcome to %v booking application.\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// userName = "Jack"
	// userTickets = 2
	// fmt.Printf("Hello %v\n", userName)
	// fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	// geting user input
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName) // & is used as a pointer. while inputing data from client (user), it must be pointing to memmory.

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)

	// displaying user input
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive the confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("The number of available remaining tickets are: %v.\n", remainingTickets)
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

}
