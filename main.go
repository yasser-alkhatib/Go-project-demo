package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These first names of all the bookings %v\n", firstNames)
			fmt.Printf("These are all the bookings %v\n", bookings)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked out. Please come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or lastname you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email is not correct, doesn not contain @")
			}
			if !isValidTicket {
				fmt.Println("number of tickets is invalid")
			}
		}
	}

	/*
		city := "London"

		switch city {
		case "New York":
			//execute code for booking New York conference tickets
		case "Singapore", "Hong Kong":
			//execute code for booking Singapore & Hong Kong conference tickets
		case "London", "Berlin":
			//some code here
		case "Mexico City":
			//some code here
		default:
			fmt.Print("No valid city is selected")
		}
	*/
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	//fmt.Printf("These first names of all the bookings %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for the user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//userData["firstName"] = firstName
	//userData["lastname"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining ticket(s): %v\n", remainingTickets)
}
