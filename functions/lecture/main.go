package main

import "fmt"

// func name(input param1 type, input param2 type) output type {
//logic or body
// }

func sum(num1, num2 int) int {
	return num1 + num2
}

func multiReturn() (int, int, int) {
	return 1, 2, 3
}

func double(x int) int {
	return x * 2
}

func greet() {
	fmt.Println("Hello there!")
}

func main() {
	greet()

	sumResult := sum(1, 2)
	fmt.Println(sumResult)

	a, b, _ := multiReturn()
	fmt.Println(a, b)

	doubleResult := double(4)
	fmt.Println(doubleResult)

	anotherDoubleResult := double(sum(double(4), 5))
	fmt.Println(anotherDoubleResult)

}
