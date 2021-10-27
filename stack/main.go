package main

import (
	"fmt"

	"github.com/iamnotjustice/ds-examples/stack/backlog"
)

func main() {
	booksBacklog := backlog.NewListBacklog()
	//booksBacklog := backlog.NewSliceBacklog()

	// add one book on top
	booksBacklog.Push(&backlog.Book{
		Author: "Peter Watts",
		Title:  "Blindsight",
	})

	fmt.Printf("Current top of the stack: %+v \n========\n", booksBacklog.Peek())

	// add another one on top
	booksBacklog.Push(&backlog.Book{
		Author: "Liu Cixin",
		Title:  "The Three-Body Problem",
	})

	fmt.Printf("Current top of the stack: %+v \n========\n", booksBacklog.Peek())

	// remove top book
	booksBacklog.Pop()

	fmt.Printf("Current top of the stack: %+v \n========\n", booksBacklog.Peek())

	// add another one on top
	booksBacklog.Push(&backlog.Book{
		Author: "Douglas Adams",
		Title:  "The Hitchhiker's Guide to the Galaxy",
	})

	backlog.PrintAll(booksBacklog)
}
