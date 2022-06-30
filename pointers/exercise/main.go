//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name        string
	securityTag SecurityTag
}

func activateSecurity(tag *SecurityTag) {
	*tag = Active
}

func deactivateSecurity(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(slice []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(slice); i++ {
		deactivateSecurity(&slice[i].securityTag)
	}
}

func main() {
	shoppingList := []Item{
		{name: "Phone", securityTag: Active},
		{name: "Tablet", securityTag: Active},
		{name: "Laptop", securityTag: Active},
		{name: "Head Phones", securityTag: Active},
	}
	fmt.Println("shopping list", shoppingList)
	deactivateSecurity(&shoppingList[1].securityTag)
	fmt.Println("tablet security deactivated", shoppingList)
	deactivateSecurity(&shoppingList[3].securityTag)
	fmt.Println("Head Phones security deactivated", shoppingList)
	checkout(shoppingList)
	fmt.Println("shopping list after checkout", shoppingList)
}
