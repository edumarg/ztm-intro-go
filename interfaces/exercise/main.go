//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftSender interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (m Car) String() string {
	return fmt.Sprintf("Car: %v", string(m))
}

func (m Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(m))
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(ls LiftSender) {
	switch ls.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", ls)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", ls)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", ls)
	}
}

func main() {
	fmt.Println("Interfaces Exercise")
	diablo := Car("Diablo")
	duccati := Motorcycle("Duccati")
	f350 := Truck("f350")

	sendToLift(diablo)
	sendToLift(duccati)
	sendToLift(f350)
}
