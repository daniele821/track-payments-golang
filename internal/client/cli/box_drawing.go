package cli

import (
	"fmt"
	"strings"
)

const boxVert string = "│"
const boxHoriz string = "─"
const boxLeftTop string = "┘"
const boxRightTop string = "└"
const boxLeftDown string = "┐"
const boxRightDown string = "┌"
const boxVertLeft string = "┤"
const boxVertRight string = "├"
const boxHorizTop string = "┴"
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
	panic("UNREACHABLE CODE!")
}

func getMaxLen(data [][]string) []int {
	maxLen := make([]int, len(data[0]))
	for _, row := range data {
		for index, cell := range row {
			if maxLen[index] < len(cell) {
				maxLen[index] = len(cell)
			}
		}
	}
	return maxLen
}

func drawBoxRow(maxLen []int, startChar, middleChar, endChar string) string {
	acc := strings.Builder{}
	for index, nthLen := range maxLen {
		if index == 0 {
			acc.WriteString(startChar)
		} else {
			acc.WriteString(middleChar)
		}
		acc.WriteString(strings.Repeat(boxHoriz, nthLen))
	}
	acc.WriteString(endChar)
	return acc.String()
}

// func fmtBox(data [][]string) {
// 	maxLen := getMaxLen()
// 	for indexRow, row := range data {
// 		if indexRow == 0 {
// 		}
// 		for indexCol, cell := range row {
//
// 		}
// 	}
// }
