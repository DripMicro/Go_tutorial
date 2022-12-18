package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("onfereneceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Welcome to %v booking application", conferenceName)
	fmt.Println("we have total of ", conferenceTickets, " tickets and", remainingTickets, " are still avialable")
	fmt.Println("Get your tickets heere to attend")

	var firsName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("enter your first name: ")
	fmt.Scan(&firsName)

	fmt.Println("enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("enter your email: ")
	fmt.Scan(&email)

	fmt.Println("enter your User Ticket: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will recieve a confirmation email at %v", firsName, lastName, userTickets, email)

}
