package component

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	introduce = `
	说明：按括号里的字母选择人物，上下左右移动，(c) 重新开始游戏.
	`
	stepStr = "step:%d"
)

var (
	// Tip prompt
	Tip string
	// CompMap is map of component
	CompMap map[int]*Component
	// PosiMap Position is exist
	PosiMap map[int8]bool

	baseX    = 10
	baseY    = 5
	baseGapX = 13
	baseGapY = 2
	baseDist = 25

	tipX    = 0
	tipY    = 2
	tipDist = 8

	// Step game step
	Step     = 0
	stepX    = 0
	stepY    = 3
	stepDist = 10

	introX = 2
	introY = 5

	initPosiDate = []int8{12, 13, 17, 18, 1, 6, 4, 9, 11, 16, 14, 19, 2, 3, 21, 7, 8, 24, 22, 23}
	initPosiSize = []int8{4, 2, 2, 2, 2, 2, 1, 1, 1, 1}
	initNameDate = []string{"曹操(q)", "张飞(w)", "黄忠(e)", "赵云(r)", "马超(t)", "关羽(y)", "小兵(u)", "小兵(i)", "小兵(o)", "小兵(p)"}

	// COLORS colors
	COLORS = []termbox.Attribute{
		termbox.ColorRed,
		termbox.ColorYellow,
		termbox.ColorWhite,
		termbox.ColorGreen,
		termbox.ColorBlue,
		termbox.ColorMagenta,
		termbox.ColorCyan,
	}
)

// Component is rectangle
type Component struct {
	Name     string
	Size     int8
	Position []int8
	BgColor  termbox.Attribute
}

// Draw himself
func (c *Component) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for _, posi := range c.Position {
		x := baseX + (int(posi)%5-1)*baseGapX
		y := baseY + int(posi)/5*baseGapY

		termbox.SetCell(x, y, ' ', termbox.ColorBlack, c.BgColor)
		termbox.Flush()
		fmt.Printf("%s", c.Name)
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// Clear himself
func (c *Component) Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for _, posi := range c.Position {
		x := baseX + (int(posi)%5-1)*13
		y := baseY + int(posi)/5*2

		termbox.SetCell(x, y, ' ', termbox.ColorBlack, termbox.ColorDefault)
		termbox.Flush()
		fmt.Print("       ")
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// Move a gap
func (c *Component) Move(gap int8) bool {
	var (
		canMoveFlag = true
		winFlag     = int8(0)
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
		Step++
		c.Clear()
		for index, posi := range c.Position {
			c.Position[index] = posi + gap
			PosiMap[posi+gap] = true

			if c.Position[index] == 22 || c.Position[index] == 23 {
				winFlag++
			}
		}

		if winFlag == 2 {
			Tip = "You are win!"
			return true
		}

		Tip = "Had move    "
		return false
	}

	for index, posi := range c.Position {
		c.Position[index] = posi
		PosiMap[posi+gap] = true
	}

	Tip = "Don't move  "

	return false
}

// DrawTip draw tip
func DrawTip() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(tipX, tipY, ' ', termbox.ColorBlack, COLORS[2])

	termbox.Flush()
	fmt.Printf("%s", Tip)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// DrawStep draw step
func DrawStep() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(stepX, stepY, ' ', COLORS[2], termbox.ColorDefault)

	termbox.Flush()
	fmt.Printf(stepStr, Step)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// DrawIntroduce draw introduce
func DrawIntroduce() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(introX, introY+5*baseGapY, ' ', termbox.ColorWhite, termbox.ColorDefault)

	termbox.Flush()
	fmt.Printf("%s", introduce)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// InitCompMap initial
func InitCompMap(file ...string) {
	var (
		posiDate  []int8
		totalSize int8
		posiIndex int8
		compIndex int
	)

	CompMap = make(map[int]*Component)
	PosiMap = make(map[int8]bool)
	tempX, _ := termbox.Size()

	tipX = tempX/2 - tipDist
	baseX = tempX/2 - baseDist
	stepX = tempX/2 + stepDist

	if len(file) == 0 {
		posiDate = initPosiDate
	} else {
		posiDate = readDateFrom(file[0])
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for ; compIndex < 10; compIndex++ {
		comp := Component{
			Name:     initNameDate[compIndex],
			Size:     initPosiSize[compIndex],
			Position: make([]int8, initPosiSize[compIndex]),
		}

		if compIndex < 6 {
			comp.BgColor = COLORS[compIndex]
		} else {
			comp.BgColor = COLORS[6]
		}

		totalSize += comp.Size
		tempIndex := 0

		for ; posiIndex < totalSize; posiIndex++ {
			comp.Position[tempIndex] = posiDate[posiIndex]
			tempIndex++
		}

		CompMap[compIndex] = &comp
		comp.Draw()
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()

	for i, posi := range posiDate {
		if i < 18 {
			PosiMap[posi] = true
		} else {
			PosiMap[posi] = false
		}
	}

	// fmt.Printf("\n[DEBUG]:CompMap %v\nPosiMap %v\n", CompMap, PosiMap)
}

func readDateFrom(file string) []int8 {
	return nil
}
