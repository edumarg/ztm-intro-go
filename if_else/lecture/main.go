package main

import (
	"fmt"
)

func average(a, b, c int) int {
	return (a + b + c) / 3
}

func main() {
	n := -3
	if n < 0 {
		fmt.Println(n, "is less than 0")
	} else if n > 0 {
		fmt.Println(n, "is bigger than 0")
	} else {
		fmt.Println("is cero")
	}

	lowerRange := 0
	upperRange := 100
	x := 101
	if x >= lowerRange && x <= upperRange {
		fmt.Println("In range")
	} else {
		fmt.Println("Out of range")
	}

	quiz1, quiz2, quiz3 := 8, 10, 5

	if quiz1 > quiz2 {
		fmt.Println("quiz1 scored higher than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 scored higher than quiz1")
	} else {
		fmt.Println("quiz2 scored same as quiz1")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("you passed")
	} else {
		fmt.Println("you failed")
	}

}
