package main

import (
	"fmt"
)

func main() {
	var example = 1
	var example2 int = 2
	var example3 int // default values: 0 for numbers, "" for strings, nil for the rest
	example3 = 3
	example4 := 4
	fmt.Println(example, example2, example3, example4) // 1 2 3 4

	var a, b, c, d = 9, 8, 7, "test"
	fmt.Println(a, b, c, d) // 9 8 7 test

	e, f := 7, "variable"
	fmt.Println(e, f) // 7 variable

	var (
		x int = 10
		y int = 20
		z     = "sample"
	)
	fmt.Println(x, y, z) // 10 20 sample

	m := 1
	n := m
	var p = n
	m = 2
	n = 3
	fmt.Println(m, n, p) // 2 3 1
}
