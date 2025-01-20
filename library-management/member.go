package main

import "fmt"

type Member struct {
	Id              int
	Name            string
	ContactInfo     string
	CurrentBorrowed []*BookItem
	BorrowHistory   []*BookItem
}

func NewMember(id int, name, contactInfo string) *Member {
	return &Member{Id: id, Name: name, ContactInfo: contactInfo,
		CurrentBorrowed: make([]*BookItem, 0), BorrowHistory: make([]*BookItem, 0)}
}

func (m *Member) isQuotaFull() bool {
	return len(m.CurrentBorrowed) >= 3
}

func (m *Member) AddBorrowedBook(bookItem *BookItem) {
	m.CurrentBorrowed = append(m.CurrentBorrowed, bookItem)
}

func (m *Member) RemoveBorrowedBook(bookItem *BookItem) {

	for i, b1 := range m.CurrentBorrowed {
		if b1.Id == bookItem.Id {
			m.CurrentBorrowed = append(m.CurrentBorrowed[:i], m.CurrentBorrowed[i+1:]...)
			break
		}
	}
	m.BorrowHistory = append(m.BorrowHistory, bookItem)
	fmt.Printf("Book %d has been added to borrow history for member %s\n",
		bookItem.Id, m.Name)
}

func (m *Member) DisplayCurentBorrowedBooks() {

	fmt.Printf("Current Borrowed books for member %s\n", m.Name)
	for _, book := range m.CurrentBorrowed {
		fmt.Printf("  %d %d\n", book.Id, book.BookId)
	}
}

func (m *Member) DisplayBorrowedHistory() {

	fmt.Printf("Current Borrowed books for member %s\n", m.Name)
	for _, book := range m.BorrowHistory {
		fmt.Printf(" Book %d(Item Id : %d) \n", book.Id, book.BookId)
	}
}
