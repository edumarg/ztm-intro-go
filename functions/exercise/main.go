//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

func greeting(name string) {
	fmt.Println("Hello", name)
}

func myMessage(message string) string {
	return message
}

func sumThreeValues(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func randNum() int {
	return rand.Intn(100)
}

func randTwoNum() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func main() {
	greeting("John")

	fmt.Println(myMessage("This is my message"))

	sumResult := sumThreeValues(1, 4, 5)
	fmt.Println(sumResult)

	num1 := randNum()
	fmt.Println(num1)

	num2, num3 := randTwoNum()
	fmt.Println(num2, num3)

	result := randNum() + sumThreeValues(1, 6, 7) + num3
	fmt.Println(result)
}
