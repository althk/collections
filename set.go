package collections

import "math"

// Set provides a collection with no duplicates.
type Set[T comparable] struct {
	elems []T
}

// Add adds the given element to the set.
// If the element already exists, it is a no-op.
func (s *Set[T]) Add(e T) {
	if !s.Contains(e) {
		s.elems = append(s.elems, e)
	}
}

// Items returns the elements of the set as a slice.
func (s *Set[T]) Items() []T {
	return s.elems
}

// Union adds all elements from s2 to the current set
// which are not present in the current set.
func (s *Set[T]) Union(s2 *Set[T]) {
	if s2 == nil {
		return
	}
	o1, o2 := bySize(s, s2)

	for _, e := range o1.Items() {
		o2.Add(e)
	}
}

// Intersection returns a new set with elements that are present in
// both the sets.
func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	if s2 == nil {
		return nil
	}
	o1, o2 := bySize(s, s2)

	common := make([]T, 0)
	for _, e := range o1.Items() {
		if o2.Contains(e) {
			common = append(common, e)
		}
	}
	return &Set[T]{elems: common}
}

// Difference returns a new set with elements that are in the current set but not in s2.
func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	if s2 == nil || s2.Size() == 0 {
		var elems []T
		copy(elems, s.elems)
		return &Set[T]{elems: elems}
	}

	diff := make([]T, 0)

	for _, e := range s.elems {
		if s2.Contains(e) {
			continue
		}
		diff = append(diff, e)
	}
	return &Set[T]{elems: diff}
}

// IsDisjoint returns true if there are no common elements between
// this set and the given set, else returns true.
func (s *Set[T]) IsDisjoint(s2 *Set[T]) bool {
	if s2 == nil || s2.Size() == 0 {
		return true
	}

	for _, e := range s2.elems {
		if s.Contains(e) {
			return false
		}
	}
	return true
}

// Contains returns true if the given element
// is already present in the current set, otherwise returns false.
func (s *Set[T]) Contains(e T) bool {
	return s.Index(e) > -1
}

// Index returns the index of the given element in the set.
// Returns -1 if the element is not present.
func (s *Set[T]) Index(e T) int {
	for i := range s.elems {
		if e == s.elems[i] {
			return i
		}
	}
	return -1
}

// IsSubsetOf returns true if this set is a subset of the given set.
func (s *Set[T]) IsSubsetOf(s2 *Set[T]) bool {
	if s2 == nil {
		return false
	}
	if s.Size() == 0 {
		return true
	}
	if s.Size() > s2.Size() {
		return false
	}

	for _, e := range s.elems {
		if !s2.Contains(e) {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the current set.
func (s *Set[T]) Size() int {
	return len(s.elems)
}

// NewSet creates and returns a new set with the given initial capacity.
func NewSet[T comparable](initialSize int) *Set[T] {
	initialSize = int(math.Max(float64(initialSize), 0))
	return &Set[T]{
		elems: make([]T, 0, initialSize),
	}
}

// FromArray creates a new set from the given array/slice.
func FromArray[T comparable](a []T) *Set[T] {
	s := NewSet[T](len(a))
	for _, e := range a {
		s.Add(e)
	}
	return s
}

// bySize returns the given sets in ascending order of their sizes.
func bySize[T comparable](s1, s2 *Set[T]) (*Set[T], *Set[T]) {
	if s1.Size() <= s2.Size() {
		return s1, s2
	}
	return s2, s1
}
