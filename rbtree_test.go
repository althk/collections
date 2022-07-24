package collections

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var rbt *RBTree[int, string]

func TestRBTree_Put(t *testing.T) {
	setup()
	require.EqualValues(t, 1000, rbt.Size())
}

func TestRBTree_GetExistingKey(t *testing.T) {
	setup()
	got, ok := rbt.Get(100)

	require.True(t, ok)
	require.EqualValues(t, "val_100", got)
}

func TestRBTree_GetNonExistentKey(t *testing.T) {
	setup()
	got, ok := rbt.Get(928342)

	require.False(t, ok)
	require.Empty(t, got)
}

func TestRBTree_Min(t *testing.T) {
	setup()
	got := rbt.Min()

	require.NotNil(t, got)
	require.EqualValues(t, 0, got.Key())
}

func TestRBTree_Max(t *testing.T) {
	setup()
	got := rbt.Max()

	require.NotNil(t, got)
	require.EqualValues(t, 999, got.Key())
}

func TestRBTree_DeleteMin(t *testing.T) {
	setup()
	rbt.DeleteMin()

	require.False(t, rbt.Has(0))
	require.EqualValues(t, 1, rbt.Min().Key())
}

func TestRBTree_DeleteMax(t *testing.T) {
	setup()
	rbt.DeleteMax()

	require.False(t, rbt.Has(999))
	require.EqualValues(t, 998, rbt.Max().Key())
}

func TestRBTree_Delete(t *testing.T) {
	setup()
	require.True(t, rbt.Has(500))

	rbt.Delete(500)

	require.False(t, rbt.Has(500))
	require.EqualValues(t, 999, rbt.Size())
}

func TestRBTree_Keys(t *testing.T) {
	setup()
	keys := rbt.Keys()
	for i := 0; i < 100; i++ {
		require.EqualValues(t, i, keys[i])
	}
}

func TestRBTree_Rank(t *testing.T) {
	setup()

	got := rbt.Rank(100)
	require.EqualValues(t, 100, got)
}

func setup() {
	arr := rand.Perm(1000)
	rbt = New[int, string]()
	for _, i := range arr {
		rbt.Put(i, fmt.Sprintf("val_%d", i))
	}
}

func TestMain(m *testing.M) {
	rand.Seed(42)
	os.Exit(m.Run())
}
