package utils_test

import (
	"payment/internal/utils"
	"testing"
)

func TestHasDuplicates(t *testing.T) {
	tests := []struct {
		name   string
		args   []any
		expect bool
	}{
		{
			name:   "HasDuplicatesInt",
			args:   []any{1, 2, 3, 3},
			expect: true,
		},
		{
			name:   "HasNoDuplicatesInt",
			args:   []any{1, 2, 3, 4},
			expect: false,
		},
		{
			name:   "HasNoDuplicatesString",
			args:   []any{"", "first", "second"},
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := utils.HasDuplicates(test.args)
			if actual != test.expect {
				t.Fatalf("%v -> expected: %t, actual: %t", test.args, test.expect, actual)
			}
		})
	}
}

func TestSortedSlice(t *testing.T) {

}
