//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printLine(title string, slice []Part) {
	fmt.Println(title)
	for i := 0; i < len(slice); i++ {
		part := slice[i]
		fmt.Println(part)
	}
}

func main() {
	fmt.Println("Slices exercise")
	line := []Part{"1", "2", "3"}
	printLine("line", line)
	line = append(line, "4", "5")
	printLine("new line", line)
	lastTwoItems := line[len(line)-2:]
	printLine("lastTwoItems", lastTwoItems)
	line = append(line, "6", "7")
	printLine("new line", line)
	lastTwoItems = line[len(line)-2:]
	printLine("lastTwoItems", lastTwoItems)
}
