package component

import (
	"github.com/nsf/termbox-go"
)

const (
	introduce = `
	说明：按括号里的字母选择人物，上下左右箭头移动。
		
		( c )   重新开始游戏
		( s )   存档
		(Esc) 退出游戏
	`
	stepStr = "step:%v"
)

var (
	// BakFile store
	BakFile = "Date.bak"
	// Tip prompt
	Tip string
	// Elems is map of component
	Elems []*Elem
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

	introX = 2
	introY = 5 + 5*baseGapY

	initPosiDate = []int8{2, 3, 7, 8, 1, 6, 4, 9, 11, 16, 14, 19, 17, 18, 21, 12, 13, 24, 22, 23}
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
