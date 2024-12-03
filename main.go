package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}
	// var bookings = []string{}
	// var bookings []string
	// var bookings [50]string
	// var bookings = [50]string{"Nana", "Nicole", "Peter"}
	// bookings[0] = "Nana"
	// bookings[1] = "Nicole"

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets now , %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		// bookings[0] = firstName + " " + lastName

		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remainig for %v. \n", remainingTickets, conferenceName)

		firstNames := []string{}
		for index, booking := range bookings {
			strings.Fields(booking)
		}

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}

}
