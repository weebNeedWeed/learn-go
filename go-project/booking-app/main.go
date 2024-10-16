package main

import (
	// "booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var MyVar = 1 // global variable M is used for exporting this variable(instead of m)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = []map[string]string{} only works for list of string, does not work for map
// var bookings = make([]map[string]string, 0) // work for map and struct
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

func main() {
	// fmt.Printf("%T", remainingTickets)
	// var bookings = [50]string{"hello", "world"}

	greetUsers()

	// for remainingTickets > 0 && len(bookings) < 50 {

	firstName, lastName, email, userTickets := getUserInput()
	// isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)
	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1) //if 2 go, use wg.Add(2)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		fmt.Printf("The value of bookings is: %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}

	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short.")
		}

		if !isValidEmail {
			fmt.Println("Email address you entered does not contain @ sign.")
		}

		if !isValidTicketNumber {
			fmt.Println("Ticket number you entered is invalid")
		}
	}
	wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName // For array

	// userData := make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // For slice

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first element: %v\n", bookings[0])
	// fmt.Printf("The type: %T\n", bookings)
	// fmt.Printf("The length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############")
	wg.Done() // decrement the counter by 1
}
