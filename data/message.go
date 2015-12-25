package data

import (
	"fmt"
	"time"
)

// Message defines properties of a single message within a convesation
type Message struct {
	// Message author
	User *User

	// Message time
	Time time.Time

	// Message text
	Text string
}

// Converts Message into string
func (m *Message) String() string {
	if m.User == nil {
		return fmt.Sprintf("user=nil time='%s' text='%s'", m.Time, m.Text)
	}

	return fmt.Sprintf("user='%s' time='%s' text='%s'", m.User.Name, m.Time, m.Text)
}
