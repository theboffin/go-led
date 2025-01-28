package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	view            *tview.Box
	app             *tview.Application
	message         = "   HELLO   "
	ledOn           = "[red:black:b]●[-:-:-]"
	ledOff          = "[gray:black:d]○[-:-:-]"
	messageLedLines []string
)

func main() {
	app = tview.NewApplication().SetInputCapture(handleKeyboard)
	view = tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrDim).
		SetTitle("[green:white] [::b]LED MATRIX ")

	view.SetDrawFunc(drawFrame)
	messageLedLines = getLedLines(message, ledOn, ledOff)

	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}

func handleKeyboard(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEscape:
		app.Stop()
	}
	return event
}

func drawFrame(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {

	for i, line := range messageLedLines {
		tview.Print(screen, line, 1, i+1, width, tview.AlignLeft, tcell.ColorWhite)
	}
	return 0, 0, 0, 0
}
