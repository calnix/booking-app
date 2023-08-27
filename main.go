package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//get user input: ask user for details
	firstName, lastName, email, userTickets := getUserInput()

	// bool values
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidTicketNumber && isValidEmail && isValidName {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//call getFirstnames
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("No more tickets!\n")
			// END THE FOR LOOP

		}
	} else {

		if !isValidName {
			fmt.Println("Name too short!")
		}

		if !isValidEmail {
			fmt.Println("Email is invalid!")
		}

		if !isValidTicketNumber {
			fmt.Println("Tickaet amount is invalid!")
		}

		fmt.Printf("Input data invalid, try again\n")

	}

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Total tickets: %v. Remaining tickets: %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!\n")
}

// bookings is a slice of strings. so define parameter as -> <param_name> []string
// now we are iterating over a slice of maps, bookings:
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// if map
		//firstNames = append(firstNames, booking["firstName"])
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

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName, lastName string, email string) uint {
	remainingTickets = remainingTickets - userTickets

	// create a map for each user
	// create a struct for each user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//	userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // 10 represents base 10 decimal number system.

	// store first name and last name in booking slice
	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thanks you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("remaining tickets %v\n", remainingTickets)

	return remainingTickets
}

func sendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //store print statement into variable
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########")

	wg.Done()
}
