package gui

import (
	"github.com/gizak/termui"
	"log"
)

var refreshGui = make(chan bool, 1)

func RefreshGUI() {
    refreshGui <- true
}

func RunGUI() {
	err := termui.Init()
	if err != nil {
		log.Printf("alert: gui   -> termui failed to init error='%s'", err)
		panic(err)
	}
	//defer termui.Close()

	height := termui.TermHeight() - 5

	msgTitle := termui.NewPar("%CHAT_TITLE%")
	msgTitle.Height = 3

	listTitle := termui.NewPar("%CHAT_LIST%")
	listTitle.Height = 3

	msgList := termui.NewList()
	msgList.Height = height
	msgList.Items = []string{
		"[13:38:00] @alex:",
		"           Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit involuptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
    msgList.Overflow = "wrap"

	chatList := NewChatList()
	chatList.SetHeight(height)
	chatList.AddOrUpdateItem("@alex 1", ClNormal)
	chatList.AddOrUpdateItem("@alex 2", ClUnread)
	chatList.AddOrUpdateItem("@alex 3", ClSelected)
	chatList.AddOrUpdateItem("#general 4", ClNormal)
	chatList.AddOrUpdateItem("#general 5", ClUnread)
	chatList.AddOrUpdateItem("#general 6", ClSelected)
	chatList.AddOrUpdateItem("@alex 7", ClNormal)
	chatList.AddOrUpdateItem("@alex 8", ClUnread)
	chatList.AddOrUpdateItem("@alex 9", ClSelected)
	chatList.AddOrUpdateItem("#general 10", ClNormal)
	chatList.AddOrUpdateItem("#general 11", ClUnread)
	chatList.AddOrUpdateItem("#general 12", ClSelected)

	inputBox := NewInputBox("> ", "Input a message and press <Enter> to send it")

	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(8, 0, msgTitle),
			termui.NewCol(4, 0, listTitle)),
		termui.NewRow(
			termui.NewCol(8, 0, msgList),
			termui.NewCol(4, 0, chatList)),
		termui.NewRow(
			termui.NewCol(12, 0, inputBox)),
		termui.NewRow(
			termui.NewCol(1, 0, NewStatusBarButton("F1", "Help")),
			termui.NewCol(1, 0, NewStatusBarButton("F2", "Switch chat")),
			termui.NewCol(2, 0, nil),
			termui.NewCol(1, 0, NewStatusBarButton("F5", "Refresh")),
			termui.NewCol(1, 0, nil),
			termui.NewCol(1, 0, NewStatusBarButton("F7", "Search")),
			termui.NewCol(1, 0, nil),
			termui.NewCol(1, 0, nil),
			termui.NewCol(1, 0, NewStatusBarButton("F10", "Exit")),
			termui.NewCol(1, 0, nil),
			termui.NewCol(1, 0, nil)))

	termui.Body.Align()
	termui.Render(termui.Body)
	//return

	for {
		select {
		case e := <-termui.EventCh():
			if e.Type == termui.EventKey {
				switch e.Key {
				case termui.KeyF10:
					return
				}
			}
        case <- refreshGui:
            termui.Body.Align()
	        termui.Render(termui.Body)
		}
	}
}

func makeStBarItem(label string) *termui.Par {
	item := termui.NewPar(label)
	item.Height = 3
	item.BgColor = termui.ColorBlue
	item.TextBgColor = termui.ColorBlue
	item.Border.BgColor = termui.ColorBlue
	return item
}
