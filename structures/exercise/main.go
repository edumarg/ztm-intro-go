//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

type Rectangle struct {
	Width, Length int
}

func rectanglePerimeter(r Rectangle) int {
	return r.Length*2 + r.Width*2
}

func rectangleArea(r Rectangle) int {
	return r.Length * r.Width
}
func printRectangleInfo(r Rectangle) {
	fmt.Println("Rectangle  Area", rectangleArea(r))
	fmt.Println("Rectangle Perimeter", rectanglePerimeter(r))
}

func main() {
	fmt.Println("Structures exercise")
	rectangle1 := Rectangle{
		Width:  3,
		Length: 6,
	}
	printRectangleInfo(rectangle1)

	var rectangle2 Rectangle
	rectangle2.Width = rectangle1.Width * 2
	rectangle2.Length = rectangle1.Length * 2
	printRectangleInfo(rectangle2)

}
