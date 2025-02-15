package cli

import (
	"fmt"
	"strings"
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
	box  int
	row  int
	cell int
}

func newCellBox(nthRow int) cell {
	return cell{box: nthRow, row: -1, cell: -1}
}
func newCellRow(nthBox, nthRow int) cell {
	return cell{box: nthBox, row: nthRow, cell: -1}
}
func newCellCol(nthBox, nthRow, nthCol int) cell {
	return cell{box: nthBox, row: nthRow, cell: nthCol}
}

func alignStr(str string, maxLen int, align align) string {
	lenStr := len(str)
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
				if maxLen[index] < len(cell) {
					maxLen[index] = len(cell)
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

func fmtBox(data [][][]string, lPad, rPad int, alignCells map[cell]align) string {
	acc := strings.Builder{}
	maxLen := getMaxLen(data)
	acc.WriteString(drawBoxRow(maxLen, boxRightDown, boxHorizDown, boxLeftDown, lPad+rPad))
	acc.WriteString("\n")
	for indexBox, box := range data {
		if indexBox != 0 {
			acc.WriteString(drawBoxRow(maxLen, boxVertRight, boxCross, boxVertLeft, lPad+rPad))
			acc.WriteString("\n")
		}
		for indexRow, row := range box {
			for indexCol, cell := range row {
				if indexCol == 0 {
					acc.WriteString(boxVert)
				}
				acc.WriteString(strings.Repeat(" ", lPad))
				align, exist := alignCells[newCellCol(indexBox, indexRow, indexCol)]
				if !exist {
					align, exist = alignCells[newCellRow(indexBox, indexRow)]
					if !exist {
						align, _ = alignCells[newCellBox(indexBox)]
					}
				}
				acc.WriteString(alignStr(cell, maxLen[indexCol], align))
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
