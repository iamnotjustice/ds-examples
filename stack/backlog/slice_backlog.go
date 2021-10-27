package backlog

import "fmt"

type SliceBacklog struct {
	books []*Book
}

func NewSliceBacklog() Backlog {
	return &SliceBacklog{
		books: make([]*Book, 0, 10),
	}
}

// Push adds new Book to the top of the backlog.
func (b *SliceBacklog) Push(newBook *Book) {
	b.books = append(b.books, newBook)
}

// Pop removes the top value of the stack.
// If stack is empty returns an error.
func (b *SliceBacklog) Pop() error {
	if len(b.books) == 0 {
		return fmt.Errorf("Backlog is empty")
	}

	b.books = b.books[:len(b.books)-1]
	return nil
}

// Peek returns top value of the stack or nil if stack is empty.
func (b *SliceBacklog) Peek() *Book {
	if len(b.books) == 0 {
		return nil
	}

	return b.books[len(b.books)-1]
}
