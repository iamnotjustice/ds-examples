package waiting_room

import "fmt"

type SliceWaitingRoom struct {
	users []*UserAccount
}

func NewSliceWaitingRoom() WaitingRoom {
	return &SliceWaitingRoom{
		users: make([]*UserAccount, 0, 10),
	}
}

// Dequeue adds one user to the line.
func (wr *SliceWaitingRoom) Enqueue(u *UserAccount) {
	wr.users = append(wr.users, u)
}

// Dequeue removes one user from the line.
func (wr *SliceWaitingRoom) Dequeue() error {
	if len(wr.users) == 0 {
		return fmt.Errorf("WaitingRoom is empty")
	}

	// trim slice to remove the first element
	wr.users = wr.users[1:]

	return nil
}

// Peek returns a user which is in front of the line (meaning he's the first to be Dequeue'd).
func (wr *SliceWaitingRoom) Peek() *UserAccount {
	if len(wr.users) == 0 {
		return nil
	}

	return wr.users[0]
}
