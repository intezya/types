package typez

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

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.data[value]
	return ok
}

func (s *Set[T]) Add(value T) {
	s.data[value] = true
}

func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, s.Size())

	for k := range s.data {
		values = append(values, k)
	}

	return values
}

func (s *Set[T]) Union(other *Set[T]) {
	for v := range other.data {
		s.data[v] = true
	}
}

func (s *Set[T]) Intersect(other *Set[T]) {
	for v := range other.data {
		if !s.data[v] {
			delete(other.data, v)
		}
	}
}

func (s *Set[T]) Clone() *Set[T] {
	clone := NewSet[T]()
	for k := range s.data {
		clone.data[k] = true
	}
	return clone
}

func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	for v := range s.data {
		if !other.data[v] {
			return false
		}
	}

	return true
}

func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	for v := range s.data {
		if !other.data[v] {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSupersetOf(other *Set[T]) bool {
	return other.IsSubsetOf(s)
}
