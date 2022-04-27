package main

import "fmt"

var showName = "Go Lang show"

func main() {

	// var showName = "🌟Go Lang show🌟"
	// const totalTickets = 50
	var remainingTickets uint = 50

	// var bookings []string
	// OR
	bookings := []string{}

	for {

		var firstName string
		var ticketsBooked uint

		// fmt.Println(hello)
		fmt.Println("🌟 Welcome to the online ticket booking for", showName, "🌟")
		fmt.Println(remainingTickets, " Tickets are available")
		fmt.Println("To book tickets, please follow the below process")
		fmt.Print("Enter your name :")
		fmt.Scan(&firstName)
		fmt.Print("Enter number of tickets :")
		fmt.Scan(&ticketsBooked)

		// This adds firstname entered by user into the bookings(array)
		bookings = append(bookings, firstName)
		remainingTickets = remainingTickets - ticketsBooked
		fmt.Println("Name :", firstName, "\nTickets booked:", ticketsBooked, "\nRemaing tickets :", remainingTickets)

		fmt.Printf("These are all our bookings => %v\n", bookings)

		fmt.Println("Thanks for choosing us,", firstName, "! Have a Great day :)")
	}

}

// var hello = "It is a hello string"
