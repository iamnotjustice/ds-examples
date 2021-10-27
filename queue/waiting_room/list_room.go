package waiting_room

import (
	"container/list"
	"fmt"
)

type ListWaitingRoom struct {
	users *list.List
}

func NewListWaitingRoom() WaitingRoom {
	return &ListWaitingRoom{
		users: list.New(),
	}
}

// Dequeue adds one user to the line.
func (wr *ListWaitingRoom) Enqueue(u *UserAccount) {
	wr.users.PushBack(u)
}

// Dequeue removes one user from the line.
func (wr *ListWaitingRoom) Dequeue() error {
	if wr.users.Len() == 0 {
		return fmt.Errorf("WaitingRoom is empty")
	}

	toDequeue := wr.users.Front()
	wr.users.Remove(toDequeue)

	return nil
}

// Peek returns a user which is in front of the line (meaning he's the first to be Dequeue'd).
func (wr *ListWaitingRoom) Peek() *UserAccount {
	if wr.users.Len() == 0 {
		return nil
	}

	u := wr.users.Front()
	if u == nil {
		return nil
	}

	if val, ok := u.Value.(*UserAccount); ok {
		return val
	}

	return nil
}
