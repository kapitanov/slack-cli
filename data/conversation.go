package data

import "fmt"

// ConversationType defines a type of a conversation
type ConversationType int

const (
	_ ConversationType = iota

	// CtUser mean a direct chat with a user
	CtUser

	// CtChannel mean a channel chat
	CtChannel

	// CtGroup mean a group chat
	CtGroup
)

// ConversationState defines a state of a conversation
type ConversationState int

const (
	_ ConversationState = iota

	// CsOK means a normal state of a conversation
	CsOK

	// CsNotFetched means an initial state of a conversation.
	// Data fetch is required.
	CsNotFetched

	// CsFetching means that conversation data is being fetched.
	// A change notification will be raised upon completion
	CsFetching
)

// ConversationEventHandler is a callback for conversation events
type ConversationEventHandler func(conversation *Conversation)

// ConversationMessageEventHandler is a callback for conversation events
type ConversationMessageEventHandler func(conversation *Conversation, message *Message)

// Conversation defines methods to work with a single conversation
type Conversation interface {
	// Gets a name of conversation
	Name() string

	// Gets a count of unread messages within a conversation
	UnreadCount() int

	// Gets a type of conversation
	Type() ConversationType

	// Gets current state of conversation
	State() ConversationState

	// Fetches messages
	Fetch()

	// Gets an array of fetched messages
	GetMessages() []*Message

	// Sends a message
	Send(message string) error

	// Adds an event handler for conversation property change event
	OnChanged(callback ConversationEventHandler)

	// Adds an event handler for new message event
	OnNewMessage(callback ConversationMessageEventHandler)
}

type conversationImpl struct {
	mgr          *conversationMgrImpl
	name         string
	unreadCount  int
	convType     ConversationType
	state        ConversationState
	onChanged    []ConversationEventHandler
	onNewMessage []ConversationMessageEventHandler
}

func newUserConversation(mgr *conversationMgrImpl, user *User) *conversationImpl {
	return &conversationImpl{
		mgr:          mgr,
		name:         user.Name,
		convType:     CtUser,
		state:        CsNotFetched,
		onChanged:    make([]ConversationEventHandler, 0),
		onNewMessage: make([]ConversationMessageEventHandler, 0),
	}
}

func newGroupConversation(mgr *conversationMgrImpl, group *Group) *conversationImpl {
	return &conversationImpl{
		mgr:          mgr,
		name:         group.Name,
		convType:     CtGroup,
		state:        CsNotFetched,
		onChanged:    make([]ConversationEventHandler, 0),
		onNewMessage: make([]ConversationMessageEventHandler, 0),
	}
}

func newChannelConversation(mgr *conversationMgrImpl, channel *Channel) *conversationImpl {
	return &conversationImpl{
		mgr:          mgr,
		name:         channel.Name,
		convType:     CtChannel,
		state:        CsNotFetched,
		onChanged:    make([]ConversationEventHandler, 0),
		onNewMessage: make([]ConversationMessageEventHandler, 0),
	}
}

func (c *conversationImpl) String() string {
	return fmt.Sprintf("type=%d name='%s'", c.convType, c.name)
}

// Gets a name of conversation
func (c *conversationImpl) Name() string {
	return c.name
}

// Gets a count of unread messages within a conversation
func (c *conversationImpl) UnreadCount() int {
	return c.unreadCount
}

// Gets a type of conversation
func (c *conversationImpl) Type() ConversationType {
	return c.convType
}

// Gets current state of conversation
func (c *conversationImpl) State() ConversationState {
	return c.state
}

// Fetches messages
func (c *conversationImpl) Fetch() {
	// TODO
}

// Gets an array of fetched messages
func (c *conversationImpl) GetMessages() []*Message {
	// TODO
	return make([]*Message, 0)
}

// Sends a message
func (c *conversationImpl) Send(message string) error {
	// TODO
	return nil
}

// Adds an event handler for conversation property change event
func (c *conversationImpl) OnChanged(callback ConversationEventHandler) {
	c.onChanged = append(c.onChanged, callback)
}

// Adds an event handler for new message event
func (c *conversationImpl) OnNewMessage(callback ConversationMessageEventHandler) {
	c.onNewMessage = append(c.onNewMessage, callback)
}
