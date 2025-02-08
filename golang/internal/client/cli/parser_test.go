package cli_test

import (
	"payment/internal/client/cli"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		expect     cli.Flags
		shouldFail bool
	}{
		{
			name:   "TestParserEmpty",
			args:   []string{},
			expect: cli.Flags{"--": ""},
		},
		{
			name:   "TestParserVarious",
			args:   []string{"--date", "2025-01-02", "--debug"},
			expect: cli.Flags{"--date": "2025-01-02", "--debug": "", "--": ""},
		},
		{
			name:   "TestParserInitVals",
			args:   []string{"2025-01-02"},
			expect: cli.Flags{"--": "2025-01-02"},
		},
		{
			name:       "TestParserFailure",
			args:       []string{"--abc", "--abc"},
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := cli.Parser(test.args)
			if test.shouldFail {
				if err == nil {
					t.Fatalf("parser should have failed!\n")
				}
			} else {
				if err != nil {
					t.Fatalf("parser should not have failed (%s)!\n", err)
				}
				if !reflect.DeepEqual(actual, test.expect) {
					t.Fatalf("parser result differs from the expect value (actual: %s, expect: %s)!", actual, test.expect)
				}
			}
		})
	}
}
