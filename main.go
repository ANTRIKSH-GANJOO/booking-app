package main

import "fmt"

func main() {

	var conferencename = "Go Conference"

	const total = 50

	var remaining = 50

	var bookings [50]string

	fmt.Println("Hello", conferencename)
	fmt.Println("We have total of", total, "tickets")
	fmt.Println("Remaining tickets", remaining)

	var username string

	var usertickets int
	fmt.Printf("Enter your Name:")
	fmt.Scan(&username)
	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&usertickets)

	fmt.Println(username)
	bookings[0] = username
	fmt.Printf("the whole array: %v\n", bookings)
	fmt.Printf("the first value: %v\n", bookings[0])
	fmt.Println(usertickets)

	remaining = remaining - usertickets

	fmt.Printf("%v tickets remaining for %v\n", remaining, conferencename)

}
