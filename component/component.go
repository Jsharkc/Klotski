package component

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// Elem is rectangle
type Elem struct {
	Name     string
	Size     int8
	Position []int8
	BgColor  termbox.Attribute
}

// Draw himself
func (c *Elem) Draw() {
	c.drawComp(c.Name, termbox.ColorBlack, c.BgColor)
}

// Clear himself
func (c *Elem) Clear() {
	c.drawComp("       ", termbox.ColorBlack, termbox.ColorDefault)
}

func (c *Elem) drawComp(cont string, fg, bg termbox.Attribute) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for _, posi := range c.Position {
		x := baseX + (int(posi)%5-1)*baseGapX
		y := baseY + int(posi)/5*baseGapY

		termbox.SetCell(x, y, ' ', fg, bg)
		termbox.Flush()
		fmt.Print(cont)
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// Move a gap
func (c *Elem) Move(gap int8) bool {
	var (
		canMoveFlag = true
	)

	for _, posi := range c.Position {
		PosiMap[posi] = false
	}

	for _, posi := range c.Position {
		exist, ok := PosiMap[posi+gap]
		if ok {
			if exist {
				canMoveFlag = false
				break
			}
			continue
		}

		canMoveFlag = false
		break
	}

	if canMoveFlag {
		c.Clear()
		for index, posi := range c.Position {
			c.Position[index] = posi + gap
			PosiMap[posi+gap] = true
		}

		if len(c.Position) == 4 && c.Position[2] == 22 && c.Position[3] == 23 {
			Tip = " You are win! "
			return true
		}

		Tip = "  Had move    "
		return false
	}

	for index, posi := range c.Position {
		c.Position[index] = posi
		PosiMap[posi+gap] = true
	}

	Tip = "  Don't move  "

	return false
}
