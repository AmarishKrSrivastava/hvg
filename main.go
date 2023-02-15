package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferencename = "Go Conference"
	const conferencetickets = 50
	var remainingtickets = 50
	var bookings []string

	fmt.Printf("Welcome to the Site %v  ", conferencename)
	fmt.Println()
	fmt.Println("We have total no  of tickets ", conferencetickets, "and we have left with", remainingtickets)
	fmt.Println("Get your tickets here")
	for {
		var Firstname string
		var lastname string
		var userticket int
		var email string

		fmt.Println("Enter the First name")
		fmt.Scan(&Firstname)
		fmt.Println("Enter the last name")
		fmt.Scan(&lastname)
		fmt.Println("Enter the email id of the user")
		fmt.Scan(&email)
		fmt.Println("Enter the no of ticket")
		fmt.Scan(&userticket)

		remainingtickets = remainingtickets - userticket
		bookings = append(bookings, Firstname+" "+lastname)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("The  slice type : %T\n", bookings)
		fmt.Printf(" slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", Firstname, lastname, userticket, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingtickets, conferencename)

		Firstnames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)

			Firstnames = append(Firstnames, names[0])
		}
		fmt.Printf("The first names of booking are : %v]\n", Firstnames)
		fmt.Println("\n\n\n\nTo break the application kindly press Ctrl+c. Thank You !")

	}

}
