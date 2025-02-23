package cli

import (
	"reflect"
	"strings"
	"testing"
)

func TestAlign(t *testing.T) {
	expected := "London   "
	actual := alignStr("London", 9, leftAlign)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("align failed: expected |%s|, actual |%s|", expected, actual)
	}

	expected = " London  "
	actual = alignStr("London", 9, centerLeftAlign)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("align failed: expected |%s|, actual |%s|", expected, actual)
	}

	expected = "  London "
	actual = alignStr("London", 9, centerRightAlign)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("align failed: expected |%s|, actual |%s|", expected, actual)
	}

	expected = "   London"
	actual = alignStr("London", 9, rightAlign)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("align failed: expected |%s|, actual |%s|", expected, actual)
	}
}

func TestBoxRow(t *testing.T) {
	s := boxRightDown
	m := boxHorizDown
	e := boxLeftDown
	expected := s + strings.Repeat(boxHoriz, 10) + m + strings.Repeat(boxHoriz, 5) + e
	actual := drawBoxRow([]int{9, 4}, s, m, e, 1)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("draw box row failed: \nexpected | %s |\nactual   | %s |", expected, actual)
	}
}

func TestBoxRow2(t *testing.T) {
	s := boxVert
	m := boxVertRight
	e := boxVertLeft
	c := boxCross
	expected := s + strings.Repeat(" ", 10) + m + strings.Repeat(boxHoriz, 5) + c + strings.Repeat(boxHoriz, 4) + e
	actual := drawBoxRow2([]int{9, 4, 3}, []bool{true, false, false}, 1)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("draw box row 2 failed: \nexpected | %s |\nactual   | %s |", expected, actual)
	}
}
