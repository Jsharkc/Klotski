package component

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nsf/termbox-go"
)

func drawStr(format string, cont interface{}, x, y int, fg, bg termbox.Attribute) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(x, y, ' ', fg, bg)

	termbox.Flush()
	fmt.Printf(format, cont)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// DrawTip draw tip
func DrawTip() {
	drawStr("%v", Tip, tipX, tipY, termbox.ColorBlack, COLORS[2])
}

// DrawIntroduce draw introduce
func DrawIntroduce() {
	drawStr("%v", introduce, introX, introY, termbox.ColorWhite, termbox.ColorDefault)
}

// InitElems initial
func InitElems(file ...string) {
	var (
		posiDate  []int8
		totalSize int8
		posiIndex int8
		elemIndex int
		err       error
	)

	Elems = make([]*Elem, 10)
	PosiMap = make(map[int8]bool)
	tempX, _ := termbox.Size()

	tipX = tempX/2 - tipDist
	baseX = tempX/2 - baseDist

	if len(file) == 0 {
		posiDate = initPosiDate
	} else {
		posiDate, err = readDateFrom(file[0])
		if err != nil {
			Tip = "Unable to load"
			DrawTip()

			posiDate = initPosiDate
		}
	}

	for ; elemIndex < 10; elemIndex++ {
		elem := Elem{
			Name:     initNameDate[elemIndex],
			Size:     initPosiSize[elemIndex],
			Position: make([]int8, initPosiSize[elemIndex]),
		}

		if elemIndex < 6 {
			elem.BgColor = COLORS[elemIndex]
		} else {
			elem.BgColor = COLORS[6]
		}

		totalSize += elem.Size
		tempIndex := 0

		for ; posiIndex < totalSize; posiIndex++ {
			elem.Position[tempIndex] = posiDate[posiIndex]
			tempIndex++
		}

		Elems[elemIndex] = &elem
		elem.Draw()
	}

	for i, posi := range posiDate {
		if i < 18 {
			PosiMap[posi] = true
		} else {
			PosiMap[posi] = false
		}
	}
}

func readDateFrom(file string) ([]int8, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	cont, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	intCont := make([]int8, 20)

	for index := range cont {
		intCont[index] = int8(cont[index])
	}

	return intCont, nil
}

// WriteDateTo write to file
func WriteDateTo(file string) error {
	tempIndex := 0
	date := make([]byte, 20)

	for index := range Elems {
		for _, posi := range Elems[index].Position {
			date[tempIndex] = byte(posi)
			tempIndex++
		}
	}

	for index := range PosiMap {
		if !PosiMap[index] {
			date[tempIndex] = byte(index)
			tempIndex++
		}
	}

	return ioutil.WriteFile(file, date, os.ModePerm)
}
