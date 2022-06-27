package main

import "fmt"

func main() {
	// Basic loops
	// for <initialization>; <condition> ; post statement {
	//logic
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---using continue---")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// While loops
	fmt.Println("---while loop---")
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// Infinite loops
	// for {
	// logic
	// if somethingHappens{
	// break
	// }
	// }
}
