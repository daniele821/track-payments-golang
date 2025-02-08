package utils_test

import (
	"payment/internal/utils"
	"testing"
)

func TestHasDuplicates(t *testing.T) {
	testSlice := []int{1, 2, 3, 3}
	if !utils.HasDuplicates(testSlice) {
		t.Fatalf("the slice has a duplicate: %v", testSlice)
	}

	testSlice = []int{1, 2, 3, 4}
	if utils.HasDuplicates(testSlice) {
		t.Fatalf("the slice has no duplicate: %v", testSlice)
	}
}
