package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println(bookings)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticket")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	bookings[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining ticket(s): %v\n", remainingTickets)

}
