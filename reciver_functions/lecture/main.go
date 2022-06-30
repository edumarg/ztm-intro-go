package main

import "fmt"

type Coordinates struct {
	X, Y int
}

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

// receiver function with pointer
func (c *Coordinates) shiftBy(x, y int) {
	c.X += x
	c.Y += y
}

// receiver function with value
func (c Coordinates) Dist(other Coordinates) Coordinates {
	return Coordinates{c.X - other.X, c.Y - other.Y}
}

func (p *ParkingLot) occupySpace(spaceNum int) {
	p.spaces[spaceNum-1].occupied = true
}

func (p *ParkingLot) emptySpace(spaceNum int) {
	p.spaces[spaceNum-1].occupied = false
}

func main() {
	coord := Coordinates{5, 5}
	coord2 := Coordinates{10, 8}
	fmt.Println(coord)
	coord.shiftBy(3, 2)
	fmt.Println("after shift", coord)
	distance := coord.Dist(coord2)
	fmt.Println("Distance between coord and coord2", distance)

	fmt.Println("===Parking Lot===")
	myParkingLot := ParkingLot{spaces: make([]Space, 10)}
	fmt.Println(myParkingLot)
	myParkingLot.occupySpace(3)
	fmt.Println(myParkingLot)
	myParkingLot.occupySpace(1)
	fmt.Println(myParkingLot)
	myParkingLot.emptySpace(3)
	fmt.Println(myParkingLot)

}
