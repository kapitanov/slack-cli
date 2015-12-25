package data

import "github.com/nlopes/slack"

// ConversationMgr defines methods to work with conversations
type ConversationMgr interface {
	// Assigns a Slack connection
	SetConnection(slackAPI *slack.Client, slackRtm *slack.RTM)

	// Gets a list of all conversations
	All() []Conversation

	// Gets or open a conversation with user/group/channel
	// Parameter arg might be either User or Channel or Group
	Get(arg interface{}) Conversation
}

// Conversations is a global instance of ConversationMgr
var Conversations ConversationMgr = &conversationMgrImpl{
	conversations:        make([]Conversation, 0),
	userConversations:    make(map[*User]*conversationImpl),
	channelConversations: make(map[*Channel]*conversationImpl),
	groupConversations:   make(map[*Group]*conversationImpl),
}

type conversationMgrImpl struct {
	conversations        []Conversation
	userConversations    map[*User]*conversationImpl
	channelConversations map[*Channel]*conversationImpl
	groupConversations   map[*Group]*conversationImpl
	slackAPI             *slack.Client
	slackRtm             *slack.RTM
}

// Assigns a Slack connection
func (mgr *conversationMgrImpl) SetConnection(slackAPI *slack.Client, slackRtm *slack.RTM) {
	mgr.slackAPI = slackAPI
	mgr.slackRtm = slackRtm
}

// Gets a list of all conversations
func (mgr *conversationMgrImpl) All() []Conversation {
	return mgr.conversations
}

// Gets or open a conversation with user/group/channel
// Parameter arg might be either User or Channel or Group
func (mgr *conversationMgrImpl) Get(arg interface{}) Conversation {
	switch u := arg.(type) {

	case *User:
		conversation, exists := mgr.userConversations[u]
		if exists {
			return conversation
		}

		userConversation := newUserConversation(mgr, u)
		mgr.userConversations[u] = userConversation
		mgr.conversations = append(mgr.conversations, userConversation)
		return userConversation

	case *Channel:
		conversation, exists := mgr.channelConversations[u]
		if exists {
			return conversation
		}

		channelConversation := newChannelConversation(mgr, u)
		mgr.channelConversations[u] = channelConversation
		mgr.conversations = append(mgr.conversations, channelConversation)
		return channelConversation

	case *Group:
		conversation, exists := mgr.groupConversations[u]
		if exists {
			return conversation
		}

		groupConversation := newGroupConversation(mgr, u)
		mgr.groupConversations[u] = groupConversation
		mgr.conversations = append(mgr.conversations, groupConversation)
		return groupConversation

	default:
		return nil
	}
}
