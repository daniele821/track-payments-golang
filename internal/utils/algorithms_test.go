package utils_test

import (
	"payment/internal/utils"
	"testing"
)

func TestHasDuplicates(t *testing.T) {
	t.Run("TestHasDuplicates1", func(t *testing.T) {
		value := []int{1, 7, 3, 4, 7}
		expected := true
		actual := utils.HasDuplicates(value)
		if expected != actual {
			t.Fatalf("test failed: expected %v, actual %v", expected, actual)
		}
	})
	t.Run("TestHasDuplicates2", func(t *testing.T) {
		value := []int{1, 1, 2, 3, 4}
		expected := false
		actual := utils.HasDuplicatesMapped(value, func(index, elem int) int { return elem*0 + index })
		if expected != actual {
			t.Fatalf("test failed: expected %v, actual %v", expected, actual)
		}
	})
}
