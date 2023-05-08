// Made by Mohamad
// Made on May 8
//
// This program tells you whether or not you can go to the museum.

package main

import (
	"fmt"
)

func main() {
	// Declare Variables
	var age int
	var weekday string

	// Ask for age
	fmt.Println("What is your age?")
	fmt.Scanln(&age)

	// Check if age is less than 13
	if age < 13 {
		// End program if the user is younger than 13
		fmt.Println("\nYou are not allowed to go to the museum.")

		fmt.Println("\nDone.")
		return
	}

	// Ask for weekday
	fmt.Println("What day of the week is it? \n(m) Monday, (t) Tuesday, (w) Wednesday, \n(th) Thursday, (f) Friday, (s) Saturday, \n(su) Sunday")
	fmt.Scanln(&weekday)

	// Check if it is Tuesday or Thursday
	switch weekday {
	case "t":
		fmt.Println("\nYou are not allowed to go to the museum.")
	case "th":
		fmt.Println("\nThe museum is closed today.")
	default:
		fmt.Println("\nYou can go to the museum today.")
	}

	fmt.Println("\nDone.")
}
