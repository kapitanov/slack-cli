package gui

import 
(
    "time"
    "github.com/gizak/termui"
)

type chatItem struct {
    Time time.Time
    Author string
    Text string
}

type ChatView struct {
    x, y, w, h int
    scroll int    
    items []*chatItem       
}

func NewChatView() *ChatView {
    return &ChatView{        
        items: make([]*chatItem, 0),
    }
}

func (view* ChatView) AddMessage(time time.Time, author, text string) {
    // TODO
}

func (view* ChatView) Clear() {
    view.items = make([]*chatItem, 0)
    view.scroll = 0
    RefreshGUI()
}

func (view* ChatView) ScrollUp() {
    // TODO
}

func (view* ChatView) ScrollDowm() {
    // TODO
}

func (view* ChatView) GetHeight() int {
	return view.h
}

func (view* ChatView) SetHeight(h int) {
	view.h = h
}

func (view* ChatView) SetWidth(w int) {
	view.w = w
}

func (view* ChatView) SetX(x int) {
	view.x = x
}

func (view* ChatView) SetY(y int) {
	view.y = y
}

func (view* ChatView) Buffer() []termui.Point {
    // TODO
    
    // Render frame
    
    // Render messages starting from N
    
    /*
Render single message:
------------------------
HH:MM:SS @username:
  MESSAGE CONTENT
  WRAPPED INTO LINES
HH:MM:SS @username:
  MESSAGE CONTENT
  WRAPPED INTO LINES
HH:MM:SS @username:
  MESSAGE CONTENT
  WRAPPED INTO LINES
HH:MM:SS @username:
  MESSAGE CONTENT
  WRAPPED INTO LINES
HH:MM:SS @username:
  MESSAGE CONTENT SINGLE LINE
HH:MM:SS @username:
  MESSAGE CONTENT SINGLE LINE
HH:MM:SS @username: MESSAGE CONTENT
  WRAPPED INTO LINES
HH:MM:SS @username: MESSAGE CONTENT
  WRAPPED INTO LINES
HH:MM:SS @username: MESSAGE CONTENT SINGLE LINE
HH:MM:SS @username: MESSAGE CONTENT SINGLE LINE
HH:MM:SS @username: MESSAGE CONTENT
  WRAPPED INTO LINES    
----------------------
    */
    
    return make([]termui.Point, 0)
}
