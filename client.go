package main

import (
	"github.com/nlopes/slack"
	"log"
	"os"
	. "slack/data"
)

var (
	slackRtm *slack.RTM
	slackAPI *slack.Client
)

// OnConnected is a callback for ConnectedEvent
func OnConnected(e *slack.ConnectedEvent) {
	log.Printf("info: slack -> Connected user='%s'", e.Info.User.Name)

	// Add known users
	for _, u := range e.Info.Users {
		user := NewUser(u.ID, u.Name)
		Users.Add(user)
	}

	// Set current user info
	me := Users.GetByID(e.Info.User.ID)
	if me != nil {
		Users.SetCurrentUser(me)
	}

	// Add known groups
	for _, g := range e.Info.Groups {
		gr := NewGroup(g.ID, g.Name)
		Groups.Add(gr)
	}

	// Add known channels
	for _, c := range e.Info.Channels {
		ch := NewChannel(c.ID, c.Name)
		Channels.Add(ch)
	}

	// TODO kickstart the GUI
}

// OnMessage is a callback for MessageEvent
func OnMessage(e *slack.MessageEvent) {
	log.Printf("debug: data  -> message type='%s' channel='%s' user='%s' time='%s' text='%s'", e.Type, e.Channel, e.User, e.Timestamp, e.Text)
}

// OnLatencyReport is a callback for LatencyReport
func OnLatencyReport(e *slack.LatencyReport) {
	log.Printf("trace: slack -> LatencyReport latency='%s'", e.Value)
}

// OnRTMError is a callback for RTMError
func OnRTMError(e *slack.RTMError) {
	EnableConsoleLog(true)
	log.Printf("alert: slack -> RTM Error code=%d msg='%s'", e.Code, e.Msg)
	os.Exit(1)
}

// OnInvalidAuthEvent is a callback for InvalidAuthEvent
func OnInvalidAuthEvent(e *slack.InvalidAuthEvent) {
	EnableConsoleLog(true)
	log.Print("alert: slack -> Invalid Auth")
	os.Exit(1)
}
