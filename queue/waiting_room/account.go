package waiting_room

import (
	"fmt"
)

type UserAccount struct {
	Username string
}

func (u *UserAccount) String() string {
	return fmt.Sprintf("[username: %s]", u.Username)
}
