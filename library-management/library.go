package main

import (
	"fmt"
	"sync"
)

var (
	libraryInstance *Library
	once            sync.Once
)

type Library struct {
	books   map[int]*Book
	members map[int]*Member
}

func GetLibraryInstance() *Library {
	once.Do(func() {
		libraryInstance = &Library{books: make(map[int]*Book), members: make(map[int]*Member)}

	})
	return libraryInstance
}

func (library *Library) AddBook(book *Book) {
	library.books[book.ID] = book
	fmt.Printf("Book %d has been added\n", book.ID)
}

func (library *Library) RemoveBook(id int) {
	delete(library.books, id)
	fmt.Printf("Book %d has been removed\n", id)
}

func (library *Library) AddMember(member *Member) {
	library.members[member.Id] = member
	fmt.Printf("Member %d has been added\n", member.Id)
}

func (library *Library) RemoveMember(id int) {
	delete(library.members, id)
	fmt.Printf("Member %d has been Removed\n", id)
}

func (library *Library) ReturnBookByMember(memberId int, bookItemId int) {
	member := library.members[memberId]
	for _, book := range member.CurrentBorrowed {
		if book.Id == bookItemId {
			book.ReturnBook()
			member.RemoveBorrowedBook(book)
			fmt.Printf("Member %d has return book %d with Item id %d", memberId, book.BookId, bookItemId)
			return
		}
	}
}

func (library *Library) DisplayAvailBooks() {

	fmt.Println("Available Books")

	for id, book := range library.books {
		if book.isBookAvailable() {
			fmt.Printf("Book %d: %s by %s (published in %s)\n",
				id, book.Title, book.Author, book.PublishingYear)

		}
	}

}

func (library *Library) BorrowBookBymember(memberid int, bookId int) (*BookItem, error) {
	if library.members[memberid] == nil || library.books[bookId] == nil {
		return nil, fmt.Errorf("Member or Book not found")
	}

	member := library.members[memberid]
	if member.isQuotaFull() {
		return nil, fmt.Errorf("Member ^d Quota is full", memberid)
	}

	book := library.books[bookId]

	if !book.isBookAvailable() {
		return nil, fmt.Errorf("Book %d is not available", bookId)
	}
	borrowBook := book.BorrowBook()
	member.AddBorrowedBook(borrowBook)
	fmt.Print("Member %d has borrowed book with item id %d", memberid, bookId, borrowBook.Id)
	return borrowBook, nil
}
