package main

import (
	comp "Klotski/component"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	COLORS = []termbox.Attribute{
		termbox.ColorWhite,
		termbox.ColorGreen,
		termbox.ColorBlue,
		termbox.ColorCyan,
		termbox.ColorMagenta,
		termbox.ColorRed,
	}
)

func main() {
	comp.InitCompMap()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	// Cleanup on exit
	defer termbox.Close()

	// Keyboard only
	termbox.SetInputMode(termbox.InputEsc)

	// Clear empty
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()

	// Event queue
	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	// Event loop
loop:
	for {
		select {
		case ev := <-event_queue:
			switch ev.Type {
			case termbox.EventKey:
				// Exit
				if ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc {
					break loop
				}

				// Restart game (if user has already lost)
				if ev.Ch == 'r' || ev.Ch == 'R' {

				}

				switch ev.Key {
				case termbox.KeyArrowLeft:
					comp.CompMap[0].Move(1)
				case termbox.KeyArrowRight:
					comp.CompMap[0].Move(1)
				case termbox.KeyArrowUp:
					comp.CompMap[0].Move(1)
				case termbox.KeyArrowDown:
					comp.CompMap[0].Move(1)
				}

			case termbox.EventResize:

			}

		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
