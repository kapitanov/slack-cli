package data

import "fmt"

// Channel is short channel description
type Channel struct {
	// Channel's ID
	ID string
	// Channel name
	Name string
}

// Converts Channel into string
func (c *Channel) String() string {
	return fmt.Sprintf("ID='%s' name='%s'", c.ID, c.Name)
}

// NewChannel creates a new instance of Channel
func NewChannel(id, name string) *Channel {
	return &Channel{
		ID:   id,
		Name: name,
	}
}
