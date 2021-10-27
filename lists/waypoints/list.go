package waypoints

import (
	"errors"
	"fmt"
)

// Singly Linked List of Geo-points.
// Not thread-safe.
type WaypointList struct {
	head *Node
	size int
}

func New() *WaypointList {
	return &WaypointList{}
}

// When there's no "Tail" pointer, you have to iterate over the whole list to Push back.
func (list *WaypointList) PushBack(p *Point) *Node {
	defer func() {
		list.size++
	}()

	newTail := &Node{
		Value: p,
	}

	// if our list is empty then we just push it as "head"
	if list.head == nil || list.size == 0 {
		list.head = newTail
		return newTail
	}

	var e *Node
	for e = list.Front(); e.next != nil; e = e.Next() {
		// iterate until we hit the "tail" node
	}

	e.next = newTail

	return newTail
}

func (list *WaypointList) PushFront(p *Point) *Node {
	defer func() {
		list.size++
	}()

	newHead := &Node{
		Value: p,
	}

	// if there's an empty head or list size is zero
	// then we have a new head
	if list.head == nil || list.size == 0 {
		list.head = newHead
		return newHead
	}

	// swap old head with a new one
	oldHead := list.head
	list.head = newHead
	newHead.next = oldHead

	return newHead
}

// InsertAfter inserts new Node after mark Node provided.
func (list *WaypointList) InsertAfter(p *Point, mark *Node) *Node {
	var e *Node
	for e = list.Front(); e != mark && e != nil; e = e.Next() {
		// iterate until we hit the mark OR hit the end
	}

	// could not find "mark"
	if e == nil {
		return nil
	}

	newNode := &Node{
		Value: p,
	}

	list.size++

	if e.next == nil {
		e.next = newNode
		return newNode
	}

	newNode.next = e.next
	e.next = newNode

	return newNode
}

// Remove deletes a node from the list (if the list contains it).
func (list *WaypointList) Remove(n *Node) *Point {
	if list.head == nil || list.size == 0 {
		return nil
	}

	// to remove head we need to replace it with the next node
	if list.head == n {
		list.head = n.next
		n.next = nil
		list.size--

		return n.Value
	}

	var e *Node
	for e = list.Front(); e.next != n && e != nil; e = e.Next() {
		// iterate until we hit the mark OR hit the end
	}

	if e == nil {
		return nil
	}

	list.size--

	// check if we're at the "tail" node
	if n.next == nil {
		e.next = nil
		return n.Value
	}

	e.next = n.next

	return n.Value
}

func (list *WaypointList) Front() *Node {
	return list.head
}

func (list *WaypointList) Len() int {
	return list.size
}

func (list *WaypointList) PrintAll() error {
	fmt.Printf("Current route is (%d point(s)):\n", list.Len())

	if list.Front() == nil {
		return errors.New("WaypointList is empty")
	}

	position := 0
	for e := list.Front(); e != nil; e = e.Next() {
		position++
		fmt.Printf("Point #%d: %v \n", position, e.Value)
	}

	fmt.Printf("=========\n")

	return nil
}
