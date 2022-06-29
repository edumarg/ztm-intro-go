package main

import "fmt"

// must have same type of data
// each piece of data is called element
// fixed size and cannot be resized
// access items by index array[<index>]
// index starts at 0

type DiscountCard struct {
	id        uint16
	storeCode byte
	rate      float32
}

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is not clean")
		}
	}
}

func main() {
	// to create an array
	// var <name of array> [<size of the array as int]type of data
	var myIntArray [3]int
	fmt.Println("myIntArray", myIntArray)
	myStringArray := [3]string{"Hello", "world", "!"}
	fmt.Println("myStringArray", myStringArray)
	myNewIntArray := [...]int{7, 8, 9} // the dots will look at the numbers of items to assign the length
	fmt.Println("myNewIntArray", myNewIntArray)
	myNewStringArray := [4]string{"first", "second", "third"} // when an element is not created it will get a default value
	fmt.Println("myNewStringArray", myNewStringArray)

	var cards [3]DiscountCard
	fmt.Println("cards", cards)
	cards[1] = DiscountCard{id: 1, storeCode: 10, rate: 0.25}
	fmt.Println("cards", cards)

	myDiscountCard := cards[1]
	fmt.Println("my Discount Card", myDiscountCard)

	//iterate over an array
	fmt.Println("iterate over cards array")
	for i := 0; i < len(cards); i++ {
		card := cards[i]
		fmt.Println("card", card)
	}

	fmt.Println("iterate over cards array using range")
	for i, card := range cards {
		fmt.Println("card index", i, "card", card)
	}

	fmt.Println("iterate over cards array using range with no index")
	for _, card := range cards {
		fmt.Println("card", card)
	}

	bedRoom := Room{name: "Bed room", cleaned: false}
	dinningRoom := Room{name: "Dinning room", cleaned: false}
	livingRoom := Room{name: "Living room", cleaned: false}
	kitchen := Room{name: "kitchen", cleaned: false}
	myHouseRooms := [4]Room{bedRoom, dinningRoom, livingRoom, kitchen}
	checkCleanliness(myHouseRooms)

	myHouseRooms[0].cleaned = true
	checkCleanliness(myHouseRooms)

	myHouseRooms[3].cleaned = true
	checkCleanliness(myHouseRooms)
}
