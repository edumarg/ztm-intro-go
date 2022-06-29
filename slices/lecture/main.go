package main

import "fmt"

// companion types that work with arrays
// enable view into an array
// dynamic and not fixed size

func iterate(slice []int) {
	// iterate over slices
	for i := 0; i < len(slice); i++ {
		item := slice[i]
		fmt.Println(item)
	}
}

func main() {
	mySlice := []int{1, 2, 3}
	item1 := mySlice[0]
	fmt.Println("my slice", mySlice)
	fmt.Println("first element", item1)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)
	fmt.Println("my slice", mySlice)
	fmt.Println("part of my slice", mySlice[2:4])
	fmt.Println("part of my slice", mySlice[:3])
	fmt.Println("part of my slice", mySlice[2:])

	myArray := [...]int{10, 9, 8, 7}
	myNewSlice1 := myArray[:2]
	fmt.Println("my myNewSlice1", myNewSlice1)
	myNewSlice2 := myArray[1:]
	fmt.Println("my myNewSlice2", myNewSlice2)
	myNewSlice3 := myNewSlice2[2:]
	fmt.Println("my myNewSlice3", myNewSlice3)

	part1 := []int{1, 2, 3}
	part2 := []int{4, 5, 6, 7, 8}
	combined := append(part1, part2...)
	fmt.Println("combined", combined)

	// preallocate size of slice
	// useful when numner of elements are know but values are not
	slice := make([]int, 10)
	fmt.Println("preallocated slice", slice)

	iterate(part1)
	iterate(part2)

	//multidimensional slices
	board := [][]string{
		{"A1", "A2"},
		{"B1", "B2"},
	}

	fmt.Println(board[0][0])
	fmt.Println(board[1][0])
	fmt.Println(board[0][1])
	fmt.Println(board[1][1])

}
