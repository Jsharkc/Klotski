package component

import (
	"fmt"
)

var (
	// Tip prompt
	Tip string
	// CompMap is map of component
	CompMap map[int]*Component
	// PosiMap Position is exist
	PosiMap map[int8]bool

	initPosiDate = []int8{2, 3, 7, 8, 1, 6, 4, 9, 11, 16, 14, 19, 12, 13, 21, 17, 18, 24, 22, 23}
	initPosiSize = []int8{4, 2, 2, 2, 2, 2, 1, 1, 1, 1}
	initNameDate = []string{"曹操(q)", "张飞(w)", "黄忠(e)", "赵云(r)", "马超(t)", "关羽(y)", "小兵1(u)", "小兵2(i)", "小兵3(o)", "小兵4(p)"}
)

// Component is rectangle
type Component struct {
	Name     string
	Position []int8
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

		Tip = "Had move"

		return false
	}

	for index, posi := range c.Position {
		c.Position[index] = posi
		PosiMap[posi+gap] = true
	}

	Tip = "Don't move"

	return false
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

	if len(file) == 0 {
		posiDate = initPosiDate
	} else {
		posiDate = readDateFrom(file[0])
	}

	for ; compIndex < 10; compIndex++ {
		comp := Component{
			Name:     initNameDate[compIndex],
			Position: make([]int8, initPosiSize[compIndex]),
		}

		totalSize += initPosiSize[compIndex]
		tempIndex := 0

		for ; posiIndex < totalSize; posiIndex++ {
			comp.Position[tempIndex] = posiDate[posiIndex]
			tempIndex++
		}

		CompMap[compIndex] = &comp
	}

	for i, posi := range posiDate {
		if i < 18 {
			PosiMap[posi] = true
		} else {
			PosiMap[posi] = false
		}
	}

	fmt.Printf("[DEBUG]:CompMap %v\nPosiMap %v\n", CompMap, PosiMap)
}

func readDateFrom(file string) []int8 {
	return nil
}
