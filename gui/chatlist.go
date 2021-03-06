package gui

import "github.com/gizak/termui"

const chatListBg = termui.ColorBlue
const chatListFg = termui.ColorCyan | termui.AttrBold

const chatListNormalItemBg = termui.ColorBlue
const chatListNormalItemFg = termui.ColorWhite | termui.AttrBold
const chatListNormalItemPrefixFg = termui.ColorWhite | termui.AttrBold

const chatListSelectItemBg = termui.ColorCyan
const chatListSelectItemFg = termui.ColorBlack
const chatListSelectItemPrefixFg = termui.ColorYellow | termui.AttrBold

const chatListUnreadItemBg = termui.ColorBlue
const chatListUnreadItemFg = termui.ColorYellow | termui.AttrBold
const chatListUnreadItemPrefixFg = termui.ColorGreen | termui.AttrBold

const chatListNormalItemPrefix = "   "
const chatListUnreadItemPrefix = "<!>"
const chatListSelectedItemPrefix = "   "

type ChatListItemState int32

const (
	_ ChatListItemState = iota
	ClNormal
	ClSelected
	ClUnread
)

type chatListItem struct {
	name  string
	state ChatListItemState
}

type ChatList struct {
	x, y, w, h int

	items []*chatListItem
}

func (list *ChatList) GetHeight() int {
	return list.h
}

func (list *ChatList) SetHeight(h int) {
	list.h = h
}

func (list *ChatList) SetWidth(w int) {
	list.w = w
}

func (list *ChatList) SetX(x int) {
	list.x = x
}

func (list *ChatList) SetY(y int) {
	list.y = y
}

func (list *ChatList) Buffer() []termui.Point {
	points := make([]termui.Point, list.w*list.h)
	stride := list.w

	// Clear bg
	for x := 0; x < list.w; x++ {
		for y := 0; y < list.h; y++ {
			points[x+stride*y] = termui.Point{
				Ch: ' ',
				X:  list.x + x,
				Y:  list.y + y,
				Bg: chatListBg,
				Fg: chatListFg,
			}
		}
	}

	setCh := func(cx, cy int, ch rune) {
		points[cx+stride*cy].Ch = ch
	}
	setFg := func(cx, cy int, fg termui.Attribute) {
		points[cx+stride*cy].Fg = fg
	}
	setBg := func(cx, cy int, bg termui.Attribute) {
		points[cx+stride*cy].Bg = bg
	}

	// Render border
	setCh(0, 0, HORIZONTAL_DOWN)
	setCh(0, list.h-1, HORIZONTAL_UP)
	setCh(list.w-1, 0, VERTICAL_LEFT)
	setCh(list.w-1, list.h-1, BOTTOM_RIGHT)

	for x := 1; x < list.w-1; x++ {
		setCh(x, 0, HORIZONTAL_LINE)
		setCh(x, list.h-1, HORIZONTAL_LINE)
	}

	for y := 1; y < list.h-1; y++ {
		setCh(0, y, VERTICAL_LINE)
		setCh(list.w-1, y, VERTICAL_LINE)
	}

	// Render items
	for i, item := range list.items {

		y := i + 1
		if y >= list.h-1 {
			break
		}

		var prefix string
		var fg, prefixFg, bg termui.Attribute
		switch item.state {
		case ClSelected:
			prefix = chatListSelectedItemPrefix
			prefixFg = chatListSelectItemPrefixFg
			fg = chatListSelectItemFg
			bg = chatListSelectItemBg
			break
		case ClUnread:
			prefix = chatListUnreadItemPrefix
			prefixFg = chatListUnreadItemPrefixFg
			fg = chatListUnreadItemFg
			bg = chatListUnreadItemBg
			break
		default:
			prefix = chatListNormalItemPrefix
			prefixFg = chatListNormalItemPrefixFg
			fg = chatListNormalItemFg
			bg = chatListNormalItemBg
			break
		}

		x := 1
		j := 0

		for x < list.w-1 && j < len(prefix) {
			setCh(x, y, rune(prefix[j]))
			setFg(x, y, prefixFg)

			x++
			j++
		}

		x++
		if x < list.w-1 {
			setCh(x, y, ' ')
			setFg(x, y, fg)
			setBg(x, y, bg)
		}

		j = 0
		for x < list.w-1 && j < len(item.name) {
			setCh(x, y, rune(item.name[j]))
			setFg(x, y, fg)
			setBg(x, y, bg)

			x++
			j++
		}

		for x < list.w-1 {
			setCh(x, y, ' ')
			setFg(x, y, fg)
			setBg(x, y, bg)

			x++
		}
	}

	return points
}

func NewChatList() *ChatList {
	return &ChatList{items: make([]*chatListItem, 0)}
}

func (list *ChatList) AddOrUpdateItem(name string, state ChatListItemState) {
	for i := 0; i < len(list.items); i++ {
		if list.items[i].name == name {
			list.items[i].state = state
			return
		}
	}

	list.items = append(list.items, &chatListItem{name: name, state: state})
}
