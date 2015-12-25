package data

import "fmt"

// Group is short group description
type Group struct {
	// Group's ID
	ID string
	// Group name
	Name string
}

// Converts Group into string
func (g *Group) String() string {
	return fmt.Sprintf("ID='%s' name='%s'", g.ID, g.Name)
}

// NewGroup creates a new instance of Group
func NewGroup(id, name string) *Group {
	return &Group{
		ID:   id,
		Name: name,
	}
}
