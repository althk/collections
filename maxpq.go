package collections

import (
	"reflect"
	"sync"
)

// Comparator is a function that compares two objects o1 and o2
// of type T and returns an integer. The comparator is used by
// the Heap to compare custom objects (non primitives).
//
// It should return a -ve number if o1 < o2, 0 if o1 == o2 and
// +ve number if o1 > o2.
type Comparator[T comparable] func(o1, o2 T) int

// MaxPQ is an implementation of Max Heap with generics support.
//
// The implementation is not thread-safe. For a thread-safe
// implementation, use ConcurentMaxPQ.
type MaxPQ[T comparable] struct {
	items      []T
	Comparator Comparator[T]
}

// Insert inserts a new element in the collection and moves it
// to the correct position.
func (pq *MaxPQ[T]) Insert(item T) {
	pq.items = append(pq.items, item)
	pq.swim(len(pq.items) - 1)
}

// PeekMax returns the current max/head element.
func (pq *MaxPQ[T]) PeekMax() T {
	return pq.items[0]
}

// DelMax returns the current max element and deletes it
// from the collection.
func (pq *MaxPQ[T]) DelMax() T {
	item := pq.items[0]
	pq.exch(0, len(pq.items)-1)
	pq.items = pq.items[:len(pq.items)-1]
	pq.sink(0)
	return item
}

func (pq *MaxPQ[T]) swim(k int) {
	for k > 0 && pq.less(k/2, k) {
		pq.exch(k/2, k)
		k = k / 2
	}
}

func (pq *MaxPQ[T]) sink(k int) {
	var j int
	for 2*k <= len(pq.items)-1 {
		if k == 0 {
			j = 1
		} else {
			j = 2 * k
		}

		if j < len(pq.items)-1 && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

func (pq *MaxPQ[T]) exch(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *MaxPQ[T]) less(i, j int) bool {
	if pq.Comparator != nil {
		return pq.Comparator(pq.items[i], pq.items[j]) < 0
	}
	switch vi, vj := reflect.ValueOf(pq.items[i]), reflect.ValueOf(pq.items[j]); vi.Kind() {
	case reflect.String:
		return vi.String() < vj.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vi.Int() < vj.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return vi.Uint() < vj.Uint()
	case reflect.Float32, reflect.Float64:
		return vi.Float() < vj.Float()
	default:
		return false
	}
}

// ConcurrentMaxPQ is the thread-safe version of MaxPQ.
// All operations of this type are thread-safe.
type ConcurrentMaxPQ[T comparable] struct {
	*MaxPQ[T]
	mu sync.RWMutex
}

// Insert inserts the given element in the collection.
func (pq *ConcurrentMaxPQ[T]) Insert(item T) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.MaxPQ.Insert(item)
}

// PeekMax returns the current max/head element.
func (pq *ConcurrentMaxPQ[T]) PeekMax() T {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.items[0]
}

// DelMax returns the current max element and deletes it
// from the collection.
func (pq *ConcurrentMaxPQ[T]) DelMax() T {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.MaxPQ.DelMax()
}

func NewMaxPQ[T comparable](capacity uint, compareFn Comparator[T]) *MaxPQ[T] {
	return &MaxPQ[T]{
		items:      make([]T, 0, capacity),
		Comparator: compareFn,
	}
}

func NewConcurrentMaxPQ[T comparable](capacity uint, compareFn Comparator[T]) *ConcurrentMaxPQ[T] {
	return &ConcurrentMaxPQ[T]{
		MaxPQ: &MaxPQ[T]{
			items:      make([]T, 0, capacity),
			Comparator: compareFn,
		},
	}
}
