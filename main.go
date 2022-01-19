package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}
	//getting types
	fmt.Printf("ConferenceTickets is %T, RemainingTickets is %T, ConferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and there are %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter you first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter you last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter you email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v remaining, so you cant book %v Tickets\n", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets! You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			//end the program
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
	}

}
