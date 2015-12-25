package main

import (                           
	"github.com/nlopes/slack"
	"log"
	"slack/gui"
)
                                                                       
func main() {

	gui.RunGUI()
	return

	log.Print("info: starting slack")
	slackAPI := slack.New(SLACK_TOKEN)
	// slackAPI.SetDebug(true)

	slackRtm = slackAPI.NewRTM()

	go slackCallbackLoop()
	slackRtm.ManageConnection()
}

func slackCallbackLoop() {
	for {
		select {
		case msg := <-slackRtm.IncomingEvents:
			slackEventHandler(&msg)
		}
	}
}

func slackEventHandler(msg *slack.RTMEvent) {
	switch e := msg.Data.(type) {
	case *slack.HelloEvent:
		log.Printf("debug: slack -> Hello")
		break

	case *slack.ConnectedEvent:
		OnConnected(e)
		break

	case *slack.MessageEvent:
		OnMessage(e)
		break

	case *slack.PresenceChangeEvent:
		log.Printf("debug: PresenceChangeEvent type=%s user=%s presence=%s", e.Type, e.User, e.Presence)
		break

	case *slack.LatencyReport:
		OnLatencyReport(e)
		break

	case *slack.RTMError:
		OnRTMError(e)
		break

	case *slack.InvalidAuthEvent:
		OnInvalidAuthEvent(e)
		break

		// ConnectionErrorEvent
		// ConnectingEvent
		// DisconnectedEvent
		// UnmarshallingErrorEvent
		// MessageTooLongEvent
		// OutgoingErrorEvent
		// IncomingEventError
		// AckErrorEvent

	default:
		log.Printf("trace: Unknown event")
		break
	}
}
