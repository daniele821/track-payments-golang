package payments

import "github.com/google/btree"

type ReadOnlyBTree[T any] struct {
	btree *btree.BTreeG[T]
}

func (readOnlyBTree ReadOnlyBTree[T]) Len() int {
	return readOnlyBTree.btree.Len()
}

func (readOnlyBTree ReadOnlyBTree[T]) Min() (T, bool) {
	return readOnlyBTree.btree.Min()
}

func (readOnlyBTree ReadOnlyBTree[T]) Max() (T, bool) {
	return readOnlyBTree.btree.Max()
}

func (readOnlyBTree ReadOnlyBTree[T]) Get(key T) (T, bool) {
	return readOnlyBTree.btree.Get(key)
}

func (readOnlyBTree ReadOnlyBTree[T]) Ascend(iterator btree.ItemIteratorG[T]) {
	readOnlyBTree.btree.Ascend(iterator)
}

func (readOnlyBTree ReadOnlyBTree[T]) Descend(iterator btree.ItemIteratorG[T]) {
	readOnlyBTree.btree.Descend(iterator)
}

func skipFirstItem[T any](item T, index *int, iterator btree.ItemIteratorG[T]) bool {
	if *index == 0 {
		return true
	}
	*index += 1
	return iterator(item)
}

func (readOnlyBTree ReadOnlyBTree[T]) AscendRange(first, last *T, includeFirst, includeLast bool, iterator btree.ItemIteratorG[T]) {
	switch {
	case first == nil && last == nil:
		readOnlyBTree.btree.Ascend(iterator)
	case first == nil && last != nil:
		readOnlyBTree.btree.AscendLessThan(*last, iterator)
	case first != nil && last == nil:
		if includeFirst {
			readOnlyBTree.btree.AscendGreaterOrEqual(*first, iterator)
		} else {
			index := 0
			readOnlyBTree.btree.AscendGreaterOrEqual(*first, func(item T) bool { return skipFirstItem(item, &index, iterator) })
		}
	case first != nil && last != nil:
		if includeFirst {
			readOnlyBTree.btree.AscendRange(*first, *last, iterator)
		} else {
			index := 0
			readOnlyBTree.btree.AscendRange(*first, *last, func(item T) bool { return skipFirstItem(item, &index, iterator) })
		}
	default:
		panic("UNREACHABLE CODE: all possible cases should have already been covered!")
	}
	if includeLast && last != nil {
		if lastItem, exists := readOnlyBTree.Get(*last); exists {
			iterator(lastItem)
		}
	}
}

func (readOnlyBTree ReadOnlyBTree[T]) DescendRange(first, last *T, includeFirst, includeLast bool, iterator btree.ItemIteratorG[T]) {
	switch {
	case first == nil && last == nil:
		readOnlyBTree.btree.Descend(iterator)
	case first == nil && last != nil:
		readOnlyBTree.btree.DescendGreaterThan(*last, iterator)
	case first != nil && last == nil:
		if includeFirst {
			readOnlyBTree.btree.DescendLessOrEqual(*first, iterator)
		} else {
			index := 0
			readOnlyBTree.btree.DescendLessOrEqual(*first, func(item T) bool { return skipFirstItem(item, &index, iterator) })
		}
	case first != nil && last != nil:
		if includeFirst {
			readOnlyBTree.btree.DescendRange(*first, *last, iterator)
		} else {
			index := 0
			readOnlyBTree.btree.DescendRange(*first, *last, func(item T) bool { return skipFirstItem(item, &index, iterator) })
		}
	default:
		panic("UNREACHABLE CODE: all possible cases should have already been covered!")
	}
	if includeLast && last != nil {
		if lastItem, exists := readOnlyBTree.Get(*last); exists {
			iterator(lastItem)
		}
	}
}
