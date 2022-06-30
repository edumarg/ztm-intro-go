package main

import "fmt"

type Counter struct {
	hits int
}

func main() {
	value := 10

	// asterisk (*) when used with a type
	var valuePtr *int
	valuePtr = &value
	fmt.Println("valuePtr", valuePtr)

	// ampersand (&) when use from variable
	newValuePtr := &value
	fmt.Println("newValuePtr", newValuePtr)

	i := 1
	increment(&i)
	fmt.Println("i", i)

	counter := Counter{}
	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)
	replaceString(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)
	replaceString(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
}

// function that accepts as parameter a pointer to an integer
func increment(x *int) {
	// asterisk(*) when use with a variable will deference the pointer and provides
	//access to the actual data it points to
	*x += 1
}

func incrementCounter(c *Counter) {
	// when using struct, there is no need of using asterisk(*) to deference.
	// the . dot notation <.key> will do that for us.
	c.hits += 1
	fmt.Println("current counter count:", c.hits)
}

func replaceString(old *string, new string, c *Counter) {
	*old = new
	incrementCounter(c)
}
