//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Status bool

const (
	CheckedOut = false
	Available  = true
)

type Member struct {
	name  string
	books []string
}

type Book struct {
	name         string
	status       bool
	lastMember   string
	checkoutTime time.Time
	returnedTime time.Time
}

func checkoutBook(b *Book, m *Member) {
	b.lastMember = m.name
	b.checkoutTime = time.Now()
	b.status = CheckedOut
	m.books = append(m.books, b.name)
}

func checkingBook(b *Book, m *Member) {
	b.returnedTime = time.Now()
	b.status = Available
	m.books = removeBook(m.books, b.name)

}

func removeBook(s []string, e string) []string {
	books := []string{}
	for i, a := range s {
		if a == e {
			books = append(s[:i], s[i+1:]...)
		}
	}
	return books
}

func main() {
	books := []Book{
		{name: "Harry Potter", status: Available},
		{name: "Moby Dick", status: Available},
		{name: "100 years solitude", status: Available},
		{name: "Mouse", status: Available},
	}
	fmt.Println("Library: ", books)
	Larry := Member{name: "Larry", books: []string{}}
	Tron := Member{name: "Tron", books: []string{}}
	John := Member{name: "John", books: []string{}}
	members := []Member{Larry, Tron, John}
	fmt.Println("Members: ", members)

	fmt.Println("----Larry Checkedout Moby Dick----")
	checkoutBook(&books[1], &members[0])
	fmt.Println("books: ", books)
	fmt.Println("members: ", members)

	fmt.Println("----Larry Checkedout Harry Potter----")
	checkoutBook(&books[0], &members[0])
	fmt.Println("books: ", books)
	fmt.Println("members: ", members)

	fmt.Println("----Larry Checkedin Moby Dick----")
	checkingBook(&books[1], &members[0])
	fmt.Println("books: ", books)
	fmt.Println("members: ", members)

}
