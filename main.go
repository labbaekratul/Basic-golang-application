package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println(("Get your tickets here to attend"))

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
