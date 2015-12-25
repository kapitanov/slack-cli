package data

import "log"

// GroupEventHandler is a callback for channel manager events
type ChannelEventHandler func(channel *Channel)

// ChannelMgr defines methods and event for channel management
type ChannelMgr interface {
	// Gets an array of all known channels
	All() []*Channel

	// Gets a channel by ID
	GetByID(id string) *Channel

	// Gets a channel by name
	GetByName(name string) *Channel

	// Adds new channel
	Add(channel *Channel)

	// Removes a channel
	Remove(channel *Channel)

	// Adds an event handler for channel addition event
	OnAdded(callback ChannelEventHandler)

	// Adds an event handler for channel removal event
	OnRemoved(callback ChannelEventHandler)
}

// Channels is a global instance of ChannelMgr
var Channels ChannelMgr = &channelsMgrImpl{
	channels: make([]*Channel, 0),
	byID:     make(map[string]*Channel),
	byName:   make(map[string]*Channel),
	added:    make([]ChannelEventHandler, 0),
	removed:  make([]ChannelEventHandler, 0),
}

type channelsMgrImpl struct {
	channels []*Channel
	byID     map[string]*Channel
	byName   map[string]*Channel

	added   []ChannelEventHandler
	removed []ChannelEventHandler
}

// Gets an array of all known groups
func (mgr *channelsMgrImpl) All() []*Channel {
	return mgr.channels
}

// Gets a channel by ID
func (mgr *channelsMgrImpl) GetByID(id string) *Channel {
	channel, _ := mgr.byID[id]
	return channel
}

// Gets a channel by name
func (mgr *channelsMgrImpl) GetByName(name string) *Channel {
	channel, _ := mgr.byName[name]
	return channel
}

// Adds new channel
func (mgr *channelsMgrImpl) Add(channel *Channel) {
	// Check if channel already exists
	if _, exists := mgr.byID[channel.ID]; exists {
		log.Printf("warning: data  -> channel already exists %s", channel)
		return
	}

	// Add a channel
	mgr.channels = append(mgr.channels, channel)
	mgr.byID[channel.ID] = channel
	mgr.byName[channel.Name] = channel

	// Raise notifications
	for _, callback := range mgr.added {
		callback(channel)
	}

	log.Printf("debug: data  -> + channel %s", channel)
}

// Removes a channel
func (mgr *channelsMgrImpl) Remove(channel *Channel) {
	// Check if channel doesn't exists
	if _, exists := mgr.byID[channel.ID]; !exists {
		log.Printf("warning: data  -> channel doesn't exists %s", channel)
		return
	}

	delete(mgr.byID, channel.ID)
	delete(mgr.byName, channel.Name)

	index := -1
	for i, c := range mgr.channels {
		if c == channel {
			index = i
			break
		}
	}

	if index >= 0 {
		// Remove an element by its index
		// See https://github.com/golang/go/wiki/SliceTricks
		copy(mgr.channels[index:], mgr.channels[index+1:])
		mgr.channels[len(mgr.channels)-1] = nil
		mgr.channels = mgr.channels[:len(mgr.channels)-1]
	}

	// Raise notifications
	for _, callback := range mgr.removed {
		callback(channel)
	}

	log.Printf("debug: data  -> - channel %s", channel)
}

// Adds an event handler for channel addition event
func (mgr *channelsMgrImpl) OnAdded(callback ChannelEventHandler) {
	mgr.added = append(mgr.added, callback)
}

// Adds an event handler for channel removal event
func (mgr *channelsMgrImpl) OnRemoved(callback ChannelEventHandler) {
	mgr.removed = append(mgr.removed, callback)
}
