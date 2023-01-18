package utils

type SliceIterator[T any] struct {
	Elements []T
	value    T
	index    int
}

// Create an iterator over the slice xs
func NewSliceIterator[T any](xs []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: xs,
	}
}

// Move to next value in collection
func (iter *SliceIterator[T]) Next() bool {
	if iter.index < len(iter.Elements) {
		iter.value = iter.Elements[iter.index]
		iter.index += 1
		return true
	}

	return false
}

// Get current element
func (iter *SliceIterator[T]) Value() T {
	return iter.value
}
