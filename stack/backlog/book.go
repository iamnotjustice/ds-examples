package backlog

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b *Book) String() string {
	return fmt.Sprintf(`"%s" by %s`, b.Title, b.Author)
}
