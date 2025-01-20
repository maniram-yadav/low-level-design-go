package main

type BookStatus string

const (
	Borrowed  BookStatus = "Borrowed"
	Available BookStatus = "Available"
)

type BookItem struct {
	Id     int
	BookId int
	Status BookStatus
}

func NewBookItem(id int, bookId int) *BookItem {
	return &BookItem{Id: id, BookId: bookId, Status: Available}
}

func (bi *BookItem) BorrowBook() {
	bi.Status = Borrowed
}

func (bi *BookItem) ReturnBook() {
	bi.Status = Available
}
