package data

import "log"

// GroupEventHandler is a callback for group manager events
type GroupEventHandler func(group *Group)

// GroupMgr defines methods and event for group management
type GroupMgr interface {
	// Gets an array of all known groups
	All() []*Group

	// Gets a group by ID
	GetByID(id string) *Group

	// Gets a group by name
	GetByName(name string) *Group

	// Adds new group
	Add(group *Group)

	// Removes a group
	Remove(group *Group)

	// Adds an event handler for group addition event
	OnAdded(callback GroupEventHandler)

	// Adds an event handler for group removal event
	OnRemoved(callback GroupEventHandler)
}

// Groups is a global instance of GroupMgr
var Groups GroupMgr = &groupsMgrImpl{
	groups:  make([]*Group, 0),
	byID:    make(map[string]*Group),
	byName:  make(map[string]*Group),
	added:   make([]GroupEventHandler, 0),
	removed: make([]GroupEventHandler, 0),
}

type groupsMgrImpl struct {
	groups []*Group
	byID   map[string]*Group
	byName map[string]*Group

	added   []GroupEventHandler
	removed []GroupEventHandler
}

// Gets an array of all known groups
func (mgr *groupsMgrImpl) All() []*Group {
	return mgr.groups
}

// Gets a group by ID
func (mgr *groupsMgrImpl) GetByID(id string) *Group {
	group, _ := mgr.byID[id]
	return group
}

// Gets a group by name
func (mgr *groupsMgrImpl) GetByName(name string) *Group {
	group, _ := mgr.byName[name]
	return group
}

// Adds new group
func (mgr *groupsMgrImpl) Add(group *Group) {
	// Check if group already exists
	if _, exists := mgr.byID[group.ID]; exists {
		log.Printf("warning: data  -> group already exists %s", group)
		return
	}

	// Add a group
	mgr.groups = append(mgr.groups, group)
	mgr.byID[group.ID] = group
	mgr.byName[group.Name] = group

	// Raise notifications
	for _, callback := range mgr.added {
		callback(group)
	}

	log.Printf("debug: data  -> + group %s", group)
}

// Removes a group
func (mgr *groupsMgrImpl) Remove(group *Group) {
	// Check if group doesn't exists
	if _, exists := mgr.byID[group.ID]; !exists {
		log.Printf("warning: data  -> group doesn't exists %s", group)
		return
	}

	delete(mgr.byID, group.ID)
	delete(mgr.byName, group.Name)

	index := -1
	for i, g := range mgr.groups {
		if g == group {
			index = i
			break
		}
	}

	if index >= 0 {
		// Remove an element by its index
		// See https://github.com/golang/go/wiki/SliceTricks
		copy(mgr.groups[index:], mgr.groups[index+1:])
		mgr.groups[len(mgr.groups)-1] = nil
		mgr.groups = mgr.groups[:len(mgr.groups)-1]
	}

	// Raise notifications
	for _, callback := range mgr.removed {
		callback(group)
	}

	log.Printf("debug: data  -> - group %s", group)
}

// Adds an event handler for group addition event
func (mgr *groupsMgrImpl) OnAdded(callback GroupEventHandler) {
	mgr.added = append(mgr.added, callback)
}

// Adds an event handler for group removal event
func (mgr *groupsMgrImpl) OnRemoved(callback GroupEventHandler) {
	mgr.removed = append(mgr.removed, callback)
}
