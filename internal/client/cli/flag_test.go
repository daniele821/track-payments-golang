package cli

import (
	"reflect"
	"testing"
)

func TestFlagParse(t *testing.T) {
	args := []string{"word1", "-abc", "word2", "-abde", "--flag1", "word3", "word4"}
	expectedFlags := FlagParsed{
		flagArgs:  map[string][]string{"": {"word1"}, "-abc": {"word2"}, "-abde": {}, "--flag1": {"word3", "word4"}},
		flagOrder: []string{"-abc", "-abde", "--flag1"},
	}
	actualFlags := ParseFlags(args)
	if !reflect.DeepEqual(expectedFlags, actualFlags) {
		t.Fatalf("flag parsing was wrong: \nexpected: %s \nactual: %s", expectedFlags, actualFlags)
	}
}
