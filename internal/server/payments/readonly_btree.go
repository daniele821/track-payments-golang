package payments

import "github.com/google/btree"

type ReadOnlyBTree[T any] struct {
	btree *btree.BTreeG[T]
}

func (readOnlyBTree *ReadOnlyBTree[T]) Len() int {
	return readOnlyBTree.btree.Len()
}

func (readOnlyBTree *ReadOnlyBTree[T]) Min() (T, bool) {
	return readOnlyBTree.btree.Min()
}

func (readOnlyBTree *ReadOnlyBTree[T]) Max() (T, bool) {
	return readOnlyBTree.btree.Max()
}

func (readOnlyBTree *ReadOnlyBTree[T]) Get(key T) (T, bool) {
	return readOnlyBTree.btree.Get(key)
}

func (readOnlyBTree *ReadOnlyBTree[T]) Ascend(iterator btree.ItemIteratorG[T], first, last *T) {
	switch {
	case first == nil && last == nil:
		readOnlyBTree.btree.Ascend(iterator)
	case first == nil && last != nil:
		readOnlyBTree.btree.AscendLessThan(*last, iterator)
	case first != nil && last == nil:
		readOnlyBTree.btree.AscendGreaterOrEqual(*first, iterator)
	case first != nil && last != nil:
		readOnlyBTree.btree.AscendRange(*first, *last, iterator)
	default:
		panic("UNREACHABLE CODE: all possible cases should have already been covered!")
	}
}

func (readOnlyBTree *ReadOnlyBTree[T]) Descend(iterator btree.ItemIteratorG[T], first, last *T) {
	switch {
	case first == nil && last == nil:
		readOnlyBTree.btree.Descend(iterator)
	case first == nil && last != nil:
		readOnlyBTree.btree.DescendGreaterThan(*last, iterator)
	case first != nil && last == nil:
		readOnlyBTree.btree.DescendLessOrEqual(*first, iterator)
	case first != nil && last != nil:
		readOnlyBTree.btree.DescendRange(*first, *last, iterator)
	default:
		panic("UNREACHABLE CODE: all possible cases should have already been covered!")
	}
}
