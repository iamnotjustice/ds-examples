package waiting_room

type WaitingRoom interface {
	Enqueue(u *UserAccount)
	Dequeue() error
	Peek() *UserAccount
}
