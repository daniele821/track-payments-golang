package flags_test

import (
	"fmt"
	"payment/internal/client/cli/flags"
	"reflect"
	"testing"
)

func TestFlagParse(t *testing.T) {
	args := []string{"word1", "-abc", "word2", "-abde", "--flag1", "word3", "word4"}
	expectedFlagArgs := map[string][]string{"": {"word1"}, "-abc": {"word2"}, "-abde": {}, "--flag1": {"word3", "word4"}}
	actualFlagParsed, _ := flags.NewFlagParsed(args)
	actualFlagArgs := actualFlagParsed.FlagArgsCopy()
	if !reflect.DeepEqual(expectedFlagArgs, actualFlagArgs) {
		t.Fatalf("flagArgs parsing failed: \nexpected: %s \nactual: %s", expectedFlagArgs, actualFlagArgs)
	}
}

func TestFlagElaboration(t *testing.T) {
	flags, _ := flags.NewFlagParsed([]string{"--debug", "word1", "word2", "-ab", "word3"})
	fmt.Println(flags)
	panic("TODO: test flag elaboration")
}
