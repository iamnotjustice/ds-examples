package backlog

import "fmt"

type Backlog interface {
	Push(book *Book)
	Pop() error
	Peek() *Book
}

// PrintAll prints all backlog books in LIFO order. Note that it "Pops" items out.
func PrintAll(bl Backlog) {
	for b := bl.Peek(); b != nil; b = bl.Peek() {
		fmt.Printf("%+v\n", b)
		bl.Pop()
	}
}
