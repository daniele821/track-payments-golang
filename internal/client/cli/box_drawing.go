package cli

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const boxVert string = "│"
const boxHoriz string = "─"
const boxLeftUp string = "┘"
const boxRightUp string = "└"
const boxLeftDown string = "┐"
const boxRightDown string = "┌"
const boxVertLeft string = "┤"
const boxVertRight string = "├"
const boxHorizUp string = "┴"
const boxHorizDown string = "┬"
const boxCross string = "┼"

/*
random box, because why not?
┌──────┬──────┐
│      │      │
│      │      │
│      │      │
├──────┼──────┤
│      │      │
│      │      │
│      │      │
└──────┴──────┘
*/

type align int

const (
	leftAlign align = iota
	centerLeftAlign
	centerRightAlign
	rightAlign
)

type cell struct {
	box   int
	row   int
	col   int
	align align
}

func moreAccurateStrLen(str string) int {
	return utf8.RuneCountInString(str)
}

func alignStr(str string, maxLen int, align align) string {
	lenStr := moreAccurateStrLen(str)
	if lenStr > maxLen {
		panic(fmt.Sprintf("invalid formatting: string (%s) is longer then max lenght (%d)", str, maxLen))
	}
	diff := maxLen - lenStr
	switch align {
	case leftAlign:
		return str + strings.Repeat(" ", diff)
	case centerLeftAlign:
		return strings.Repeat(" ", diff/2) + str + strings.Repeat(" ", diff/2+diff%2)
	case centerRightAlign:
		return strings.Repeat(" ", diff/2+diff%2) + str + strings.Repeat(" ", diff/2)
	case rightAlign:
		return strings.Repeat(" ", diff) + str
	}
	panic("UNREACHABLE CODE: invalid value of align!")
}

func getMaxLen(data [][][]string) []int {
	maxLen := make([]int, len(data[0][0]))
	for _, box := range data {
		for _, row := range box {
			for index, cell := range row {
				if maxLen[index] < moreAccurateStrLen(cell) {
					maxLen[index] = moreAccurateStrLen(cell)
				}
			}
		}
	}
	return maxLen
}

func drawBoxRow(maxLen []int, startChar, middleChar, endChar string, totPad int) string {
	acc := strings.Builder{}
	for index, nthLen := range maxLen {
		if index == 0 {
			acc.WriteString(startChar)
		} else {
			acc.WriteString(middleChar)
		}
		acc.WriteString(strings.Repeat(boxHoriz, nthLen+totPad))
	}
	acc.WriteString(endChar)
	return acc.String()
}

func fmtBox(data [][][]string, lPad, rPad int, alignCells []cell) string {
	return fmtBoxVersions(data, lPad, rPad, alignCells, 0)
}

func fmtBox2(data [][][]string, lPad, rPad int, alignCells []cell) string {
	return fmtBoxVersions(data, lPad, rPad, alignCells, 1)
}

func fmtBoxVersions(data [][][]string, lPad, rPad int, alignCells []cell, version int) string {
	acc := strings.Builder{}
	maxLen := getMaxLen(data)
	acc.WriteString(drawBoxRow(maxLen, boxRightDown, boxHorizDown, boxLeftDown, lPad+rPad))
	acc.WriteString("\n")
	for indexBox, box := range data {
		if indexBox != 0 {
			switch version {
			case 0:
				acc.WriteString(drawBoxRow(maxLen, boxVertRight, boxCross, boxVertLeft, lPad+rPad))
			case 1:
				skip := []bool{}
				for _, elem := range box[0] {
					skip = append(skip, strings.TrimSpace(elem) == "")
				}
				acc.WriteString(drawBoxRow2(maxLen, skip, lPad+rPad))
			default:
				panic("invalid version")
			}
			acc.WriteString("\n")
		}
		for indexRow, row := range box {
			for indexCol, cell := range row {
				if indexCol == 0 {
					acc.WriteString(boxVert)
				}
				acc.WriteString(strings.Repeat(" ", lPad))
				alignByPreciseness := [4]*align{}
				for _, alignCell := range alignCells {
					alignBox, alignRow, alignCol := alignCell.box, alignCell.row, alignCell.col
					calc := map[int]int{-1: 1}
					preciseness := calc[alignBox] + calc[alignRow] + calc[alignCol]
					if (alignBox == -1 || alignBox == indexBox) && (alignRow == -1 || alignRow == indexRow) && (alignCol == -1 || alignCol == indexCol) {
						alignByPreciseness[preciseness] = &alignCell.align
					}
				}
				cellAlign := leftAlign
				for i := range 4 {
					if alignByPreciseness[i] != nil {
						cellAlign = *alignByPreciseness[i]
						break
					}
				}
				acc.WriteString(alignStr(cell, maxLen[indexCol], cellAlign))
				acc.WriteString(strings.Repeat(" ", rPad))
				acc.WriteString(boxVert)
			}
			acc.WriteString("\n")
		}
	}
	acc.WriteString(drawBoxRow(maxLen, boxRightUp, boxHorizUp, boxLeftUp, lPad+rPad))
	acc.WriteString("\n")
	return acc.String()
}

func drawBoxRow2(maxLen []int, skip []bool, totPad int) string {
	acc := strings.Builder{}
	left, right := false, false
	for index, nthLen := range maxLen {
		if index > 0 {
			left = !skip[index-1]
		}
		right = !skip[index]
		if left && right {
			acc.WriteString(boxCross)
		} else if left && !right {
			acc.WriteString(boxVertLeft)
		} else if !left && right {
			acc.WriteString(boxVertRight)
		} else if !left && !right {
			acc.WriteString(boxVert)
		}
		if skip[index] {
			acc.WriteString(strings.Repeat(" ", nthLen+totPad))
		} else {
			acc.WriteString(strings.Repeat(boxHoriz, nthLen+totPad))
		}
	}
	if skip[len(skip)-1] {
		acc.WriteString(boxVert)
	} else {
		acc.WriteString(boxVertLeft)
	}

	return acc.String()
}
