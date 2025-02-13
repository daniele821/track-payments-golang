package cli_test

import (
	"payment/internal/client/cli"
	"reflect"
	"testing"
)

func TestFlagParse(t *testing.T) {
	args := []string{"word1", "-abc", "word2", "-abde", "--flag1", "word3", "word4"}
	expectedFlagArgs := map[string][]string{"": {"word1"}, "-abc": {"word2"}, "-abde": {}, "--flag1": {"word3", "word4"}}
	expectedFlagOrder := []string{"-abc", "-abde", "--flag1"}
	actualFlagParsed := cli.ParseFlags(args)
	actualFlagArgs := actualFlagParsed.FlagArgsCopy()
	actualFlagOrder := actualFlagParsed.FlagOrderCopy()
	if !reflect.DeepEqual(expectedFlagArgs, actualFlagArgs) {
		t.Fatalf("flagArgs parsing failed: \nexpected: %s \nactual: %s", expectedFlagArgs, actualFlagArgs)
	}
	if !reflect.DeepEqual(expectedFlagOrder, actualFlagOrder) {
		t.Fatalf("FlagOrder parsing failed: \nexpected: %s \nactual: %s", expectedFlagOrder, actualFlagOrder)
	}
}
