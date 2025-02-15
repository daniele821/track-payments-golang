package cli

import (
	"reflect"
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
