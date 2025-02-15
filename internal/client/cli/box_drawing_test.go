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

func TestBoxDrawing(t *testing.T) {
	expected := `
┌────────┬───────────────┐
│ Index  │     City      │
├────────┼───────────────┤
│ 0      │ Tesseract     │
│   1    │     Monopoly  │
├────────┼───────────────┤
│ 2      │ VeryLongWord  │
│ 3      │ Hello!        │
└────────┴───────────────┘
`
	actual := "\n" + fmtBox([][][]string{{{"Index", "City"}},
		{{"0", "Tesseract"}, {"1", "Monopoly"}},
		{{"2", "VeryLongWord"}, {"3", "Hello!"}}},
		1, 2,
		[]cell{
			{box: 1, row: 1, col: 1, align: rightAlign},
			{box: 1, row: 1, col: -1, align: centerRightAlign},
			{box: 0, row: -1, col: -1, align: centerRightAlign},
		})
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("draw box failed: \nexpected:\n%s\nactual:\n%s\n", expected, actual)
	}
}
