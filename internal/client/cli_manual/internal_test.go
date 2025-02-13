package cli_manual

import (
	"reflect"
	"testing"
)

func TestMatchAnyLength(t *testing.T) {
	if !matchEveryLenght("cit", "city") {
		t.Fatal("fails match 1")
	}
	if !matchEveryLenght("city", "city") {
		t.Fatal("fails match 2")
	}
	if !matchEveryLenghtFromAnyWords("citi", []string{"city", "cities"}) {
		t.Fatal("fails match 3")
	}
	if matchEveryLenght("", "city") {
		t.Fatal("should fail match 4")
	}
	if matchEveryLenght("citya", "city") {
		t.Fatal("should fail match 5")
	}
	if matchEveryLenghtFromAnyWords("citi", []string{"city", "cit"}) {
		t.Fatal("should fail match 6")
	}
}

func TestSplitter(t *testing.T) {
	expected := [][]string{{"word1@", "wor@d2"}, {"", "--"}, {}}
	actual := splitter([]string{"word1@", "wor@d2", "@", "", "--", "@"})
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("splitter failed: \nactual: %s\nexpected: %s", actual, expected)
	}
}
