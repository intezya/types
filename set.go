package types

import "maps"

type Set[T comparable] struct {
	data map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]bool)}
}

func SetFromSlice[T comparable](values []T) *Set[T] {
	set := NewSet[T]()

	for i := range values {
		set.data[values[i]] = true
	}

	return set
}

func (s *Set[T]) Iter() *Iterator[T] {
	data := maps.Keys(s.data)
	return IteratorFromSeq(data)
}
