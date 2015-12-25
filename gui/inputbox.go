package gui

import "github.com/gizak/termui"

const inputBoxPromptFg = termui.ColorYellow | termui.AttrBold
const inputBoxPromptBg = termui.ColorBlack
const inputBoxValueFg = termui.ColorWhite | termui.AttrBold
const inputBoxValueBg = termui.ColorBlack

type InputBox struct {
	x, y, w       int
	prompt, value string
}

func (box *InputBox) GetHeight() int {
	return 1
}

func (box *InputBox) SetWidth(w int) {
	box.w = w
}

func (box *InputBox) SetX(x int) {
	box.x = x
}

func (box *InputBox) SetY(y int) {
	box.y = y
}

func (box *InputBox) Buffer() []termui.Point {
	points := make([]termui.Point, box.w)
	i := 0

	// Prompt
	for j := 0; j < len(box.prompt) && i < box.w; {
		points[i] = termui.Point{
			Ch: rune(box.prompt[j]),
			X:  box.x + i,
			Y:  box.y,
			Bg: inputBoxPromptBg,
			Fg: inputBoxPromptFg,
		}

		i++
		j++
	}

	// Value
	for j := 0; j < len(box.value) && i < box.w; {
		points[i] = termui.Point{
			Ch: rune(box.value[j]),
			X:  box.x + i,
			Y:  box.y,
			Bg: inputBoxValueBg,
			Fg: inputBoxValueFg,
		}

		i++
		j++
	}

	return points
}

func NewInputBox(prompt, value string) *InputBox {
	return &InputBox{prompt: prompt, value: value}
}
