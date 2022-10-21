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

type Id int
type Name string
type Title string

type Book struct {
	id     Id
	title  Title
	lended int
	total  int
}

type LendAudit struct {
	checkInTime  time.Time
	checkOutTime time.Time
}

type Member struct {
	id    Id
	name  Name
	books map[Id]LendAudit
}

type Library struct {
	books   map[Id]Book
	members map[Id]Member
}

func bookCheckOut(library *Library, bookId Id, memberId Id) bool {
	book, found := library.books[bookId]
	if !found {
		fmt.Println("Book not found.")
		return false
	}
	if book.lended == book.total {
		fmt.Println("Book is not available.")
		return false
	}
	member, found := library.members[memberId]
	if !found {
		fmt.Println("Member not found.")
		return false
	}
	book.lended += 1
	library.books[bookId] = book
	member.books[book.id] = LendAudit{checkOutTime: time.Now()}
	library.members[memberId] = member
	fmt.Println("Book checked out: ", book.title)
	return true
}

func bookCheckIn(library *Library, bookId Id, memberId Id) bool {
	book, found := library.books[bookId]
	if !found {
		fmt.Println("Book not found.")
		return false
	}
	member, found := library.members[memberId]
	if !found {
		fmt.Println("Member not found.")
		return false
	}
	audit, found := member.books[book.id]
	if !found {
		fmt.Println("Member has no book.")
		return false
	}
	book.lended -= 1
	library.books[bookId] = book
	audit.checkInTime = time.Now()
	member.books[book.id] = audit
	library.members[memberId] = member
	fmt.Println("Book checked in:", book.title)
	return true
}

func printLibrary(library *Library) {
	fmt.Println("Books:")
	for _, book := range library.books {
		fmt.Println(book)
	}
	fmt.Println("Members:")
	for _, member := range library.members {
		fmt.Println(member)
	}
}

func main() {
	books := []Book{
		{
			id:     1,
			title:  "Book 1",
			total:  1,
			lended: 0,
		},
		{
			id:     2,
			title:  "Book 2",
			total:  2,
			lended: 0,
		},
		{
			id:     3,
			title:  "Book 3",
			total:  3,
			lended: 0,
		},
		{
			id:     4,
			title:  "Book 4",
			total:  4,
			lended: 0,
		},
	}

	members := []Member{
		{
			id:    1,
			name:  "Member 1",
			books: make(map[Id]LendAudit),
		},
		{
			id:    2,
			name:  "Member 2",
			books: make(map[Id]LendAudit),
		},
		{
			id:    3,
			name:  "Member 3",
			books: make(map[Id]LendAudit),
		},
	}

	library := Library{
		books:   make(map[Id]Book),
		members: make(map[Id]Member),
	}
	for i := 0; i < len(books); i++ {
		book := books[i]
		library.books[book.id] = book
	}
	for i := 0; i < len(members); i++ {
		member := members[i]
		library.members[member.id] = member
	}
	printLibrary(&library)
	bookCheckOut(&library, 1, 1)
	printLibrary(&library)
	bookCheckIn(&library, 1, 1)
	printLibrary(&library)
}
