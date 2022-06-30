package main

import "fmt"

// variadics is use with functions to alow them to accept any number of parameters

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)

	answer := sum(c...)
	fmt.Println(answer)

}
