package collections

import (
	"reflect"
	"sync"
)

type Comparator[T comparable] func(o1, o2 T) int

type MaxPQ[T comparable] struct {
	items      []T
	Comparator Comparator[T]
}

func (pq *MaxPQ[T]) Insert(item T) {
	pq.items = append(pq.items, item)
	pq.swim(len(pq.items) - 1)
}

func (pq *MaxPQ[T]) PeekMax() T {
	return pq.items[0]
}

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

type ConcurrentMaxPQ[T comparable] struct {
	*MaxPQ[T]
	mu sync.RWMutex
}

func (pq *ConcurrentMaxPQ[T]) Insert(item T) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.MaxPQ.Insert(item)
}

func (pq *ConcurrentMaxPQ[T]) PeekMax() T {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.items[0]
}

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
