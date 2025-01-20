package main

import "fmt"

func main() {

	library := GetLibraryInstance()
	book1 := NewBook(1, "John Smith", "1997", "Learn Go")
	book2 := NewBook(2, "Bobby", "2007", "Holkn")
	book3 := NewBook(3, "J Rick", "1992", "Bikaner")

	library.AddBook(book1)
	library.AddBook(book2)
	library.AddBook(book3)

	library.DisplayAvailBooks()

	member1 := NewMember(1, "Bob proctor", "190-456-9872")
	member2 := NewMember(2, "Michelle", "237-154-39080")

	library.AddMember(member1)
	library.AddMember(member2)

	user1Borrowed1, err := libraryInstance.BorrowBookBymember(member1.Id, book1.ID)
	if err != nil {
		fmt.Println("error in borrowing book ", err)
	}

	user2Borrowed2, err := libraryInstance.BorrowBookBymember(member2.Id, book2.ID)
	if err != nil {
		fmt.Println("error in borrowing book ", err)
	}

	user2Borrowed1, err := libraryInstance.BorrowBookBymember(member2.Id, book1.ID)
	if err != nil {
		fmt.Println("error in borrowing book ", err)
	}

	member1.DisplayCurentBorrowedBooks()
	member2.DisplayCurentBorrowedBooks()

	library.DisplayAvailBooks()

	library.ReturnBookByMember(member1.Id, user1Borrowed1.BookId)
	library.ReturnBookByMember(member2.Id, user2Borrowed2.BookId)
	library.ReturnBookByMember(member2.Id, user2Borrowed1.BookId)

	member1.DisplayBorrowedHistory()
	member2.DisplayBorrowedHistory()
	library.DisplayAvailBooks()

}
