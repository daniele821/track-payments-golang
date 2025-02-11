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
