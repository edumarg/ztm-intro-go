//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	fmt.Println("variable exercises")

	var myColor string = "blue"
	fmt.Println("my favorite colo is:", myColor)

	birthYear, age := 1982, 40
	fmt.Println("I was born in", birthYear, "and I'm", age, "years old")

	var (
		firstInitial string = "E"
		lastInitial  string = "M"
	)
	fmt.Println(firstInitial + lastInitial)

	var ageDays uint32
	ageDays = uint32(age) * uint32(365)
	fmt.Println(ageDays)
}
