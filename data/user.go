package data

import "fmt"

// User is short user description
type User struct {
	// User's ID
	ID string
	// User name
	Name string
}

// Converts User into string
func (u *User) String() string {
	return fmt.Sprintf("ID='%s' name='%s'", u.ID, u.Name)
}

// NewUser creates a new instance of User
func NewUser(id, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}
