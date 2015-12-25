package data

import "log"

// UserEventHandler is a callback for user manager events
type UserEventHandler func(user *User)

// UserMgr defines methods and event for user management
type UserMgr interface {
	// Gets an array of all known users
	All() []*User

	// Gets current user
	Me() *User

	// Gets a user by ID
	GetByID(id string) *User

	// Gets a user by name
	GetByName(name string) *User

	// Adds new user
	Add(user *User)

	// Removes a user
	Remove(user *User)

	// Sets current user
	SetCurrentUser(user *User)

	// Adds an event handler for user addition event
	OnAdded(callback UserEventHandler)

	// Adds an event handler for user removal event
	OnRemoved(callback UserEventHandler)
}

// Users is a global instance of UserMgr
var Users UserMgr = &userMgrImpl{
	users:   make([]*User, 0),
	byID:    make(map[string]*User),
	byName:  make(map[string]*User),
	added:   make([]UserEventHandler, 0),
	removed: make([]UserEventHandler, 0),
}

type userMgrImpl struct {
	users  []*User
	byID   map[string]*User
	byName map[string]*User

	me *User

	added   []UserEventHandler
	removed []UserEventHandler
}

// Gets an array of all known users
func (mgr *userMgrImpl) All() []*User {
	return mgr.users
}

// Gets current user
func (mgr *userMgrImpl) Me() *User {
	return mgr.me
}

// Gets a user by ID
func (mgr *userMgrImpl) GetByID(id string) *User {
	user, _ := mgr.byID[id]
	return user
}

// Gets a user by name
func (mgr *userMgrImpl) GetByName(name string) *User {
	user, _ := mgr.byName[name]
	return user
}

// Adds new user
func (mgr *userMgrImpl) Add(user *User) {
	// Check if user already exists
	if _, exists := mgr.byID[user.ID]; exists {
		log.Printf("warning: data  -> user already exists %s", user)
		return
	}

	// Add a user
	mgr.users = append(mgr.users, user)
	mgr.byID[user.ID] = user
	mgr.byName[user.Name] = user

	// Raise notifications
	for _, callback := range mgr.added {
		callback(user)
	}

	log.Printf("debug: data  -> + user %s", user)
}

// Removes a user
func (mgr *userMgrImpl) Remove(user *User) {
	// Check if user doesn't exists
	if _, exists := mgr.byID[user.ID]; !exists {
		log.Printf("warning: data  -> user doesn't exists %s", user)
		return
	}

	delete(mgr.byID, user.ID)
	delete(mgr.byName, user.Name)

	index := -1
	for i, u := range mgr.users {
		if u == user {
			index = i
			break
		}
	}

	if index >= 0 {
		// Remove an element by its index
		// See https://github.com/golang/go/wiki/SliceTricks
		copy(mgr.users[index:], mgr.users[index+1:])
		mgr.users[len(mgr.users)-1] = nil
		mgr.users = mgr.users[:len(mgr.users)-1]
	}

	// Raise notifications
	for _, callback := range mgr.removed {
		callback(user)
	}

	log.Printf("debug: data  -> - user %s", user)
}

// Sets current user
func (mgr *userMgrImpl) SetCurrentUser(user *User) {
	// Make sure that user exists
	if mgr.GetByID(user.ID) == nil {
		mgr.Add(user)
	}

	mgr.me = user
	log.Printf("debug: data  -> me %s", user)
}

// Adds an event handler for user addition event
func (mgr *userMgrImpl) OnAdded(callback UserEventHandler) {
	mgr.added = append(mgr.added, callback)
}

// Adds an event handler for user removal event
func (mgr *userMgrImpl) OnRemoved(callback UserEventHandler) {
	mgr.removed = append(mgr.removed, callback)
}
