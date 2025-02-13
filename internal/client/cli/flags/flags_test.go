package flags_test

import (
	"fmt"
	"payment/internal/client/cli/flags"
	"reflect"
	"testing"
)

func TestFlagParse(t *testing.T) {
	args := []string{"-abc", "word2", "-de", "--flag1", "word3", "word4"}
	expectedFlagArgs := map[string][]string{"-abc": {"word2"}, "-de": {}, "--flag1": {"word3", "word4"}}
	actualFlagParsed, err := flags.NewFlagParsed(args)
	if err != nil {
		t.Fatalf("error while parsing flags: %s", err)
	}
	actualFlagArgs := actualFlagParsed.FlagArgsCopy()
	if !reflect.DeepEqual(expectedFlagArgs, actualFlagArgs) {
		t.Fatalf("flagArgs parsing failed: \nexpected: %s \nactual: %s", expectedFlagArgs, actualFlagArgs)
	}
}

func TestFlagElaboration(t *testing.T) {
	flags, _ := flags.NewFlagParsed([]string{"--debug", "word1", "word2", "-ab", "word3"})
	fmt.Println(flags, "TODO!")
}
