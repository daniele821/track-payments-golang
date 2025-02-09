package utils

import (
	"cmp"
	"errors"
	"slices"
)

type SortedSlice[T cmp.Ordered] []T

func NewSortedSlice[T cmp.Ordered](slice []T) SortedSlice[T] {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)
	slices.Sort(sortedSlice)
	return sortedSlice
}

func (sortedSlice SortedSlice[T]) Find(elem T) (int, error) {
	index, found := slices.BinarySearch(sortedSlice, elem)
	var err error = nil
	if !found {
		err = errors.New("could not find the element!")
	}
	return index, err
}

func Insert[T cmp.Ordered](elem T) {

}
