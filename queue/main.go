package main

import (
	"fmt"

	waitroom "github.com/iamnotjustice/ds-examples/queue/waiting_room"
)

func main() {
	//room := waitroom.NewListWaitingRoom()
	room := waitroom.NewSliceWaitingRoom()

	// add first in line
	room.Enqueue(&waitroom.UserAccount{
		Username: "John Doe",
	})

	// add second in line
	room.Enqueue(&waitroom.UserAccount{
		Username: "Mac DeMarco",
	})

	fmt.Printf("Next to log in => %+v\n", room.Peek())

	// add third in line
	room.Enqueue(&waitroom.UserAccount{
		Username: "Phil Spencer",
	})

	fmt.Printf("Next to log in => %+v\n", room.Peek())

	// one moved over
	room.Dequeue()

	fmt.Printf("Next to log in => %+v\n", room.Peek())
}
