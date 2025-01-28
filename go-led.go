package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	view            *tview.Box
	app             *tview.Application
	message         = "HELLO"
	ledOn           = "[red:black:b]●[-:-:-]"
	ledOff          = "[gray:black:d]○[-:-:-]"
	simpleLedOff    = "○"
	messageLedLines []string
	messageLedCount int
	displaySpeed    = 15
	messageHead     = -1
)

func main() {
	app = tview.NewApplication().SetInputCapture(handleKeyboard)
	view = tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrDim).
		SetTitle("[green:white] [::b]LED MATRIX ")

	view.SetDrawFunc(drawFrame)
	messageLedLines, messageLedCount = getLedLines(message, ledOn, ledOff)

	go refresh()

	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}

func handleKeyboard(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEscape:
		app.Stop()
	case tcell.KeyUp:
		displaySpeed = max(displaySpeed-1, 1)
	case tcell.KeyDown:
		displaySpeed = min(displaySpeed+1, 100)

	}
	return event
}

func refresh() {
	tick := time.NewTicker(time.Duration(displaySpeed) * time.Millisecond)
	for range tick.C {
		// Refresh the screen.
		app.Draw()
		// Reset the timer to the current speed.
		tick.Reset(time.Duration(displaySpeed) * time.Millisecond)
	}
}

func drawFrame(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	if messageHead == -1 {
		messageHead = width
	}

	fullRow := "[gray:black:d]" + strings.Repeat(simpleLedOff, width-2)

	for i, line := range messageLedLines {
		tview.Print(screen, fullRow, 1, i+1, width, tview.AlignLeft, tcell.ColorWhite)
		tview.Print(screen, line, messageHead, i+1, width-messageHead-1, tview.AlignLeft, tcell.ColorWhite)
	}

	msg := fmt.Sprintf("[width=%d, height=%d, Speed=%d]", width, height, displaySpeed)
	tview.Print(screen, msg, x, 2+len(messageLedLines), width, tview.AlignCenter, tcell.ColorLime)

	if messageHead > 1 {
		messageHead--
	} else {
		messageHead = width
	}

	return 0, 0, 0, 0
}
