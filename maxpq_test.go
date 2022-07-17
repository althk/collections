package collections

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type temp struct {
	val int
}

func Compare(o1, o2 temp) int {
	return o1.val - o2.val
}

func TestMaxPQ_PeekMax(t *testing.T) {
	pq := setupPQWithInts(t)

	require.EqualValues(t, 9, pq.PeekMax())
	require.EqualValues(t, 10, len(pq.items))

	pq.Insert(22)
	require.EqualValues(t, 22, pq.PeekMax())

}

func TestMaxPQ_DelMax(t *testing.T) {
	pq := setupPQWithInts(t)

	require.EqualValues(t, 9, pq.DelMax())
	require.EqualValues(t, 9, len(pq.items))

	require.EqualValues(t, 8, pq.DelMax())
	require.EqualValues(t, 8, len(pq.items))
}

func TestMaxPQStruct_DelMax(t *testing.T) {
	pq := setupPQWithStruct(t)

	require.EqualValues(t, temp{val: 9}, pq.DelMax())
	require.EqualValues(t, 9, len(pq.items))

	require.EqualValues(t, temp{val: 8}, pq.DelMax())
	require.EqualValues(t, 8, len(pq.items))
}

func setupPQWithInts(t *testing.T) *MaxPQ[int] {
	pq := NewMaxPQ[int](10, nil)

	for i := 0; i < 10; i++ {
		pq.Insert(i)
	}
	require.EqualValues(t, 10, len(pq.items))
	return pq
}

func setupPQWithStruct(t *testing.T) *MaxPQ[temp] {
	pq := NewMaxPQ(10, Compare)

	for i := 0; i < 10; i++ {
		pq.Insert(temp{
			val: i,
		})
	}
	require.EqualValues(t, 10, len(pq.items))
	return pq
}

func TestConcurrentMaxPQ_DelMax(t *testing.T) {
	pq := NewConcurrentMaxPQ[int](20, nil)
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i <= 1; i++ {
		go func(i int) {
			for j := i * 10; j < i*10+10; j++ {
				pq.Insert(j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	require.EqualValues(t, 20, len(pq.items))
	require.EqualValues(t, 19, pq.DelMax())
	require.EqualValues(t, 19, len(pq.items))

}
