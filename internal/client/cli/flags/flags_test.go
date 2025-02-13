package flags_test

import (
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
	found, words, err := flags.GetFlag("--debug", true)
	if !found {
		t.Fatalf("should have found the flag --debug")
	} else if !reflect.DeepEqual(words, []string{"word1", "word2"}) {
		t.Fatalf("invalid words: %v", words)
	} else if err != nil {
		t.Fatalf("error happened: %s", err)
	}
	found, words, err = flags.GetFlag("-a", true)
	if !found {
		t.Fatalf("should have found the flag -a")
	} else if !reflect.DeepEqual(words, []string{"word3"}) {
		t.Fatalf("invalid words: %v", words)
	} else if err != nil {
		t.Fatalf("error happened: %s", err)
	}
	found, words, err = flags.GetFlag("-b", false)
	if !found {
		t.Fatalf("should have found the flag -b")
	} else if !reflect.DeepEqual(words, []string{}) {
		t.Fatalf("invalid words: %v", words)
	} else if err != nil {
		t.Fatalf("error happened: %s", err)
	}
	found, words, err = flags.GetFlag("-c", false)
	if found {
		t.Fatalf("should not have found the flag -c")
	} else if !reflect.DeepEqual(words, []string{}) {
		t.Fatalf("invalid words: %v", words)
	} else if err == nil {
		t.Fatalf("error should have happened")
	}
	err = flags.Conclude()
	if err != nil {
		t.Fatalf("conclusion failed: %s", err)
	}
}
