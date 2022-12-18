package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to ", conferenceName, " booking application")
	fmt.Println("we have total of ", conferenceTickets, " tickets and", remainingTickets, " are still avialable")
	fmt.Println("Get your tickets heere to attend")

}
