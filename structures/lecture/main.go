package main

import "fmt"

// type <structure name> struct {
// <field name> <field type>
//}

type Sample struct {
	field string
	a, b  int
}

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	data := Sample{"word", 1, 2}
	fmt.Println("data", data)
	data.field = "new word"
	fmt.Println("data", data)

	data2 := Sample{
		field: "test",
		a:     3,
		b:     4,
	}
	fmt.Println("data2", data2)
	fmt.Println("data2.a", data2.a)

	data3 := Sample{}
	fmt.Println("data3", data3)
	data3.b = 5
	fmt.Println("data3", data3)

	// anonymous structures
	var sample struct { // inline structs created using var will have default values
		field string
		a, b  int
	}
	sample.field = "hello"
	fmt.Println("sample", sample)

	sample2 := struct { // with shorthand you must have the fields assign to values
		field string
		a, b  int
	}{"Hi!!",
		1,
		2,
	}
	fmt.Println("sample2", sample2)

	// example
	casey := Passenger{
		"Casey",
		1,
		false,
	}
	fmt.Println("casey", casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println("bill and ella", bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println("Heidi", heidi)

	casey.Boarded = true
	if casey.Boarded {
		fmt.Println("Casey has boarded")
	}

	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has boarded")
	}

	heidi.Boarded = true
	if heidi.Boarded {
		fmt.Println("Heidi has boarded")
	}

	bus := Bus{FrontSeat: heidi}
	fmt.Println("bus", bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

}
