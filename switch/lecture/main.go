package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	x := 5
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("other: ", x)
	}

	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap")
	case p < 10:
		fmt.Println("Moderately")
	default:
		fmt.Println("Expensive")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy")
	case Business:
		fmt.Println("Business")
	case FirstClass:
		fmt.Println("FirstClass")
	default:
		fmt.Println("other")
	}

	switch role, today := Guest, Monday; {
	case role == Manager:
		accessGranted()
	case role == Admin:
		accessGranted()
	case role == Contractor && (today == Saturday || today == Sunday):
		accessGranted()
	case role == Member && (today >= Monday || today <= Friday):
		accessGranted()
	case role == Guest && (today == Monday || today == Wednesday || today == Friday):
		accessGranted()
	default:
		accessDenied()

	}

}
