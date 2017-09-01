package main

import (
	"flag"
	"time"

	"github.com/nsf/termbox-go"

	comp "Klotski/component"
)

var (
	curComp *comp.Elem
	isWin   bool
)

// DrawAll draw game UI
func DrawAll() {
	for _, value := range comp.Elems {
		value.Draw()
	}
	comp.DrawTip()
}

// ClearAll clear game UI
func ClearAll() {
	for _, value := range comp.Elems {
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

	flag.Parse()
	if len(flag.Args()) > 0 {
		comp.InitElems(flag.Arg(0))
	} else {
		comp.InitElems()
	}

	comp.DrawIntroduce()
	curComp = comp.Elems[0]

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
					curComp = comp.Elems[0]
				case 'w', 'W':
					curComp = comp.Elems[1]
				case 'e', 'E':
					curComp = comp.Elems[2]
				case 'r', 'R':
					curComp = comp.Elems[3]
				case 't', 'T':
					curComp = comp.Elems[4]
				case 'y', 'Y':
					curComp = comp.Elems[5]
				case 'u', 'U':
					curComp = comp.Elems[6]
				case 'i', 'I':
					curComp = comp.Elems[7]
				case 'o', 'O':
					curComp = comp.Elems[8]
				case 'p', 'P':
					curComp = comp.Elems[9]
				case 'c', 'C':
					ClearAll()
					comp.InitElems()
					isWin = false
					curComp = comp.Elems[0]
				case 's', 'S':
					err := comp.WriteDateTo(comp.BakFile)
					if err != nil {
						comp.Tip = " Store Error  "
					}
					comp.Tip = "Store Success "
				default:
				}

				comp.DrawTip()

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
