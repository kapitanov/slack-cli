package gui

import "github.com/gizak/termui"

const statusBarButtonKeyWidth = 2
const statusBarKeyFg = termui.ColorWhite
const statusBarKeyBg = termui.ColorBlack
const statusBarLabelFg = termui.ColorBlack
const statusBarLabelBg = termui.ColorCyan

type statusBarButton struct {
	key, label string
	x, y, w    int
}

func (btn *statusBarButton) GetHeight() int {
	return 1
}

func (btn *statusBarButton) SetWidth(w int) {
	btn.w = w
}

func (btn *statusBarButton) SetX(x int) {
	btn.x = x
}

func (btn *statusBarButton) SetY(y int) {
	btn.y = y
}

func (btn *statusBarButton) Buffer() []termui.Point {
	points := make([]termui.Point, btn.w)
	i := 0

	if btn.w >= statusBarButtonKeyWidth {
		// Key
		for ; i < len(btn.key); i++ {
			points[i] = termui.Point{
				Ch: rune(btn.key[i]),
				X:  btn.x + i,
				Y:  btn.y,
				Bg: statusBarKeyBg,
				Fg: statusBarKeyFg,
			}
		}

		// Key placeholder fill
		for ; i < statusBarButtonKeyWidth; i++ {
			points[i] = termui.Point{
				Ch: ' ',
				X:  btn.x + i,
				Y:  btn.y,
				Bg: statusBarLabelBg,
				Fg: statusBarKeyFg,
			}
		}

		// Label
		for j := 0; i < btn.w && j < len(btn.label); {
			r := rune(btn.label[j])
			if i == btn.w-1 && j != len(btn.label)-1 {
				r = '…'
			}

			points[i] = termui.Point{
				Ch: r,
				X:  btn.x + i,
				Y:  btn.y,
				Bg: statusBarLabelBg,
				Fg: statusBarLabelFg,
			}

			i++
			j++
		}

		// Label placeholder fill
		for ; i < btn.w; i++ {
			points[i] = termui.Point{
				Ch: ' ',
				X:  btn.x + i,
				Y:  btn.y,
				Bg: termui.ColorCyan,
				Fg: statusBarLabelFg,
			}
		}
	}

	return points
}

func NewStatusBarButton(key, label string) termui.GridBufferer {
	return &statusBarButton{
		key:   key,
		label: label,
	}
}
