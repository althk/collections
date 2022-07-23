package collections

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var rbt *RBTree[string, string]

func TestRBTree_Put(t *testing.T) {
	setup()
	require.EqualValues(t, 1000, rbt.Size())
}

func TestRBTree_GetExistingKey(t *testing.T) {
	setup()
	got, ok := rbt.Get("key_100")

	require.True(t, ok)
	require.EqualValues(t, "val_100", got)
}

func TestRBTree_GetNonExistentKey(t *testing.T) {
	setup()
	got, ok := rbt.Get("nonexistent_key")

	require.False(t, ok)
	require.Empty(t, got)
}

func TestRBTree_Min(t *testing.T) {
	setup()
	got := rbt.Min()

	require.NotNil(t, got)
	require.EqualValues(t, "key_0", got.Key())
}

func TestRBTree_Max(t *testing.T) {
	setup()
	got := rbt.Max()

	require.NotNil(t, got)
	require.EqualValues(t, "key_999", got.Key())
}

func TestRBTree_DeleteMin(t *testing.T) {
	setup()
	rbt.DeleteMin()

	require.False(t, rbt.Has("key_0"))
	require.EqualValues(t, "key_1", rbt.Min().Key())
}

func TestRBTree_DeleteMax(t *testing.T) {
	setup()
	rbt.DeleteMax()

	require.False(t, rbt.Has("key_999"))
	require.EqualValues(t, "key_998", rbt.Max().Key())
}

func TestRBTree_Delete(t *testing.T) {
	setup()
	require.True(t, rbt.Has("key_500"))

	rbt.Delete("key_500")

	require.False(t, rbt.Has("key_500"))
	require.EqualValues(t, 999, rbt.Size())
}

func setup() {
	arr := rand.Perm(1000)
	rbt = New[string, string]()
	for _, i := range arr {
		rbt.Put(fmt.Sprintf("key_%d", i), fmt.Sprintf("val_%d", i))
	}
}
func TestMain(m *testing.M) {
	rand.Seed(42)
	os.Exit(m.Run())
}
