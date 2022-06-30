package main

import "fmt"

// iota are use with constants
// use to automatically assign values
// iota will start with 0 and will increment by 1 until all constant have a value

const (
	Online      = iota // 0
	Offline            // 1
	Maintenance        // 2
	Retired            // 3
)

// this is similar to writing
// const (
// 	Online = 0
// 	Offline = 1
// 	Maintenance = 2
// 	Retired = 3
// )

// skip values
const (
	s0 = iota // 0
	_
	_
	s3 //3
	s4 //4
)

// start at a specific value
const (
	i3 = iota + 3 //3
	i4            //4
	i5            //5
)

// iota enumeration pattern
type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return fmt.Sprintf("North")
	case East:
		return fmt.Sprintf("East")
	case South:
		return fmt.Sprintf("South")
	case West:
		return fmt.Sprintf("West")
	default:
		return fmt.Sprintf("no dirrection")
	}
}

func main() {
	north := North
	fmt.Println(north) // North
}
