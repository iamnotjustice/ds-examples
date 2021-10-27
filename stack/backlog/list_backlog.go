package backlog

import (
	"container/list"
	"fmt"
)

// Backlog of books implemented using Linked List.
type ListBacklog struct {
	books *list.List
}

func NewListBacklog() Backlog {
	return &ListBacklog{
		books: list.New(),
	}
}

// Push adds new Book to the top of the backlog.
func (b *ListBacklog) Push(newBook *Book) {
	b.books.PushFront(newBook)
}

// Pop removes the top value of the stack.
// If stack is empty returns an error.
func (b *ListBacklog) Pop() error {
	if b.books.Len() == 0 {
		return fmt.Errorf("Backlog is empty")
	}

	toRemove := b.books.Front()
	b.books.Remove(toRemove)

	return nil
}

// Peek returns top value of the stack or nil if stack is empty.
func (b *ListBacklog) Peek() *Book {
	frontBook := b.books.Front()
	if frontBook == nil {
		return nil
	}

	if val, ok := frontBook.Value.(*Book); ok {
		return val
	}

	return nil
}
