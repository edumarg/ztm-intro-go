package main

import (
	"errors"
	"fmt"
)

type DivError struct {
	a, b int
}

func divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by 0")
	} else {
		return num1 / num2, nil
	}
}

func (d *DivError) Error() string {
	return fmt.Sprintf("Cannot divide by zero: %d/%d", d.a, d.b)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivError{a, b}
	} else {
		return a / b, nil
	}
}

// ----

type Stuff struct {
	values []int
}

func (s *Stuff) get(index int) (int, error) {
	if index >= len(s.values) {
		return -1, errors.New(fmt.Sprintf("no element found at index: %v", index))
	} else {
		return s.values[index], nil
	}
}

func main() {
	answer1, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The answer is: ", answer1)
	}

	answer2, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The answer is: ", answer2)
	}

	answer3, err := div(3, 0)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("The answer is: ", answer3)
	}

	answer4, err := div(9, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The answer is: ", answer4)
	}

	myStuff := Stuff{
		values: []int{
			10,
			20,
			5,
			30,
			40},
	}

	value, err := myStuff.get(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is:", value)
	}

}
