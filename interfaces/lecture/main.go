package main

import "fmt"

// Interfaces allow the user to specify the behavior of a type instead of the type itself when
// creating functions, allowing them to operate on more than one type of data.

// ----- Example 1

type MyInterface interface {
	Function1()
	Function2(x int) int
}

func (m MyType) Function1() {}
func (m MyType) Function2(x int) int {
	return x + x
}

func execute(i MyInterface) {
	i.Function1()
}

// ----- Example 2

type Resetter interface {
	Reset()
}

type MyType int

type Coordinate struct {
	x int
	y int
}

type Player struct {
	health   int
	position Coordinate
}

func (p *Player) Reset() {
	p.health = 100
	p.position = Coordinate{x: 0, y: 0}
}

func Reset(r Resetter) {
	r.Reset()
}

// ----- Example 3

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	// logic to get greeting
	return "Hello!!"
}

func (sb spanishBot) getGreeting() string {
	// logic to get greeting
	return "Hola!!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// ----- Example 4
type dish interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Prepare Chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("add salad dressing")
}

func prepareDish(d dish) {
	fmt.Println("Prepare Dishes:")
	fmt.Printf("---Dish: %v---\n", d)
	d.PrepareDish()
	fmt.Println()
}

func main() {
	fmt.Println("Interfaces lecture")

	// ----- Example 1

	m := MyType(1)
	execute(m)  // by value
	execute(&m) // by reference

	// ----- Example 2

	player := Player{50, Coordinate{x: 5, y: 5}}
	fmt.Println("player", player)
	Reset(&player)
	fmt.Println("player after reset", player)

	// ----- Example 3

	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

	// ----- Example 4

	var friedChicken Chicken = "Fried Chicken"
	var cesarSalad Salad = "Cesar Salad"

	prepareDish(friedChicken)
	prepareDish(cesarSalad)

}
