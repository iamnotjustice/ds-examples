package waypoints

type Node struct {
	next *Node

	Value *Point
}

func (n *Node) Next() *Node {
	if n == nil || n.next == nil {
		return nil
	}

	return n.next
}
