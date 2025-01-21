package main

import "sync"

type Book struct {
	ID             int
	Author         string
	PublishingYear string
	Title          string
	BookItem       []BookItem
	lock           sync.Mutex
}

func NewBook(id int, author string, year string, title string) *Book {
	books := &Book{ID: id, Title: title, Author: author, PublishingYear: year,
		BookItem: make([]BookItem, 0)}

	for i := 1; i <= 10; i++ {
		books.BookItem = append(books.BookItem, *NewBookItem(i, id))
	}
	return books
}

func (b *Book) isBookAvailable() bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	for i := range b.BookItem {
		if b.BookItem[i].Status == Available {
			return true
		}
	}
	return false
}

func (b *Book) BorrowBook() *BookItem {
	b.lock.Lock()
	defer b.lock.Unlock()
	for i := range b.BookItem {
		if b.BookItem[i].Status == Available {
			b.BookItem[i].BorrowBook()
			return &b.BookItem[i]
		}
	}
	return nil
}
