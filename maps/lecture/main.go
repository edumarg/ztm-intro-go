package main

import "fmt"

// structures to store data as key key-values pair
// high performance
// unordered
// all keys must have same type
// all values must have same type

func main() {
	// <map name> := make(map[<type of the keys>]<type of the key values>)
	myMap := make(map[string]int)
	fmt.Println("myMap", myMap)

	// <map name> := map[<type of the keys>]<type of the key values>{
	//<key 1>:<key value 1>
	//<key 2>:<key value 2>
	//<key 3>:<key value 3>
	//}
	myMap2 := map[string]int{
		"item 1": 1,
		"item 2": 2,
		"item 3": 3,
	}
	fmt.Println("myMap2", myMap2)

	myMap2["item 4"] = 4
	fmt.Println("myMap2", myMap2)
	fmt.Println("myMap2 item 4", myMap2["item 4"])

	delete(myMap2, "item 4")
	fmt.Println("myMap2", myMap2)

	item3, found := myMap2["item 3"]
	if !found {
		fmt.Println("item 3 not found")
		return
	} else {
		fmt.Println("item 3:", item3)
	}

	// map iteration
	for key, value := range myMap2 {
		fmt.Println(key, ":", value)
	}

}
