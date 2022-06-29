//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func getLastItem(sl [4]Product) Product {
	var lastItem Product
	for i := 0; i < len(sl); i++ {
		item := sl[i]
		if item.name != "" {
			lastItem = item
		}
	}
	return lastItem
}

func calculateTotalPrice(sl [4]Product) float32 {
	var total float32
	for i := 0; i < len(sl); i++ {
		item := sl[i]
		if item.price != 0 {
			total += item.price
		}
	}
	return total
}

func totalItems(sl [4]Product) uint16 {
	var total uint16
	for i := 0; i < len(sl); i++ {
		item := sl[i]
		if item.name != "" {
			total += 1
		}
	}
	return total
}

func main() {
	fmt.Println("Arrays exercise")
	product1 := Product{name: "PC", price: 155.99}
	product2 := Product{name: "Tablet", price: 54.79}
	product3 := Product{name: "Phone", price: 75.00}
	myShoppingList := [4]Product{product1, product2, product3}
	fmt.Println("last item on the shopping list", getLastItem(myShoppingList))
	fmt.Println("total items on the shopping list", totalItems(myShoppingList))
	fmt.Println("total value of the shopping list", calculateTotalPrice(myShoppingList))

	product4 := Product{name: "Phone", price: 85.50}
	myShoppingList[len(myShoppingList)-1] = product4
	fmt.Println("last item on the shopping list", getLastItem(myShoppingList))
	fmt.Println("total items on the shopping list", totalItems(myShoppingList))
	fmt.Println("total value of the shopping list", calculateTotalPrice(myShoppingList))

}
