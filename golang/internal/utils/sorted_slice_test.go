package utils_test

import (
	"payment/internal/utils"
	"testing"
)

func TestSortedSlice_find(t *testing.T) {
	t.Run("SortedSliceSearch", func(t *testing.T) {
		arr := utils.NewSortedSlice([]int{1, 4, 2, 3, 5, 0})
		if index, err := arr.Find(3); err != nil {
			t.Fatalf("search failed (%s)!", err)
		} else if index != 3 {
			t.Fatalf("search returned wrong index (expected: 3, actual %d)!", index)
		}
	})
}
