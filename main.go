package main

import (
	"time"

	"github.com/nsf/termbox-go"

	comp "Klotski/component"
)

var (
	curComp *comp.Component
	isWin   bool
)

// DrawAll draw game UI
func DrawAll() {
	for _, value := range comp.CompMap {
		value.Draw()
	}
	comp.DrawTip()
	comp.DrawStep()
}

// ClearAll clear game UI
func ClearAll() {
	for _, value := range comp.CompMap {
		value.Clear()
	}
	comp.DrawTip()
}

func main() {
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

	comp.InitCompMap()
	comp.DrawIntroduce()
	curComp = comp.CompMap[0]

	// Event queue
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	// Event loop
loop:
	for {
		select {
		case ev := <-eventQueue:
			switch ev.Type {
			case termbox.EventKey:
				// Exit
				if ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc {
					break loop
				}

				// Restart game (if user has already lost)

				switch ev.Ch {
				case 'q', 'Q':
					curComp = comp.CompMap[0]
				case 'w', 'W':
					curComp = comp.CompMap[1]
				case 'e', 'E':
					curComp = comp.CompMap[2]
				case 'r', 'R':
					curComp = comp.CompMap[3]
				case 't', 'T':
					curComp = comp.CompMap[4]
				case 'y', 'Y':
					curComp = comp.CompMap[5]
				case 'u', 'U':
					curComp = comp.CompMap[6]
				case 'i', 'I':
					curComp = comp.CompMap[7]
				case 'o', 'O':
					curComp = comp.CompMap[8]
				case 'p', 'P':
					curComp = comp.CompMap[9]
				case 'c', 'C':
					ClearAll()
					comp.InitCompMap()
					isWin = false
					curComp = comp.CompMap[0]
				default:
				}

				if !isWin {
					switch ev.Key {
					case termbox.KeyArrowLeft:
						isWin = curComp.Move(-1)
					case termbox.KeyArrowRight:
						isWin = curComp.Move(1)
					case termbox.KeyArrowUp:
						isWin = curComp.Move(-5)
					case termbox.KeyArrowDown:
						isWin = curComp.Move(5)
					}

					DrawAll()
				}

			case termbox.EventResize:
			}
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
