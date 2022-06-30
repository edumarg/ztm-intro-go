package main

import "fmt"

func surround(msg string, left, right rune) string {
	return fmt.Sprintf("%c%v%c", left, msg, right)
}

func main() {
	//For terminal
	// fmt.Printf() // Custom format
	// %v default value
	// %t true or false
	// %c Character
	// %X Hex
	// %U Unicode format
	// %e Scientific notation
	value := 45
	fmt.Printf("%v\n", value) // 45
	fmt.Printf("%c\n", value) // - from ascii table
	fmt.Printf("%X\n", value) // 2D from ascii table
	fmt.Printf("%U\n", value)
	fmt.Printf("%e\n", value)

	// fmt.Print()   // Simple print
	// fmt.Println() // Simple print with new line

	// For data streams
	// fmt.Fprintf() // Prints to data stream: Custom format
	// %v default value
	// %t true or false
	// %c Character
	// %X Hex
	// %U Unicode format
	// %e Scientific notation

	// fmt.Fprint()   // Prints to data stream:Simple print
	// fmt.Fprintln() // Prints to data stream: Simple print with new line

	// Creating new strings
	// fmt.Sprintf() // Prints to new string: Custom format
	// %v default value
	// %t true or false
	// %c Character
	// %X Hex
	// %U Unicode format
	// %e Scientific notation
	surrounded := surround("This message", '(', ')')
	fmt.Println(surrounded) // (This Message)

	// fmt.Sprint()   // Prints to new string:Simple print
	// fmt.Sprintln() // Prints to new string: Simple print with new line
}
