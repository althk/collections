package collections

import "math"

type Comparable interface {
	int64 | int | float64 | uint | uint64 | string
}

type CompareFn[K Comparable] func(key1, key2 K) int

type Node[K Comparable, V any] struct {
	key   K
	val   V
	left  *Node[K, V]
	right *Node[K, V]
	size  int
	color Color
}

func (n *Node[K, V]) Key() K {
	return n.key
}

func (n *Node[K, V]) Val() V {
	return n.val
}

type RBTree[K Comparable, V any] struct {
	root      *Node[K, V]
	compareFn CompareFn[K]
}

type Color int

const (
	Red Color = iota
	Black
)

func New[K Comparable, V any]() *RBTree[K, V] {
	return &RBTree[K, V]{}
}

func NewWithComparator[K Comparable, V any](compareFn CompareFn[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		compareFn: compareFn,
	}
}

func (t *RBTree[K, V]) Put(key K, val V) {
	t.root = put(t.root, key, val, t.compareFn)
	t.root.color = Black
}

func (t *RBTree[K, V]) Get(key K) (V, bool) {
	return get(t.root, key, t.compareFn)
}

func (t *RBTree[K, V]) IsEmpty() bool {
	return t.root.size == 0
}

func (t *RBTree[K, V]) Min() *Node[K, V] {
	if t.IsEmpty() {
		return &Node[K, V]{}
	}
	return min(t.root)
}

func (t *RBTree[K, V]) Max() *Node[K, V] {
	if t.IsEmpty() {
		return &Node[K, V]{}
	}
	return max(t.root)
}

func (t *RBTree[K, V]) Size() int {
	return size(t.root)
}

func (t *RBTree[K, V]) DeleteMin() {
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = Red
	}
	t.root = deleteMin(t.root)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *RBTree[K, V]) DeleteMax() {
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = Red
	}
	t.root = deleteMax(t.root)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *RBTree[K, V]) Delete(key K) {
	if !t.Has(key) {
		return
	}

	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = Red
	}

	t.root = deleteNode(t.root, key, t.compareFn)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *RBTree[K, V]) Height() int {
	return int(height(t.root))
}

func (t *RBTree[K, V]) Has(key K) bool {
	_, ok := t.Get(key)
	return ok
}

func (t *RBTree[K, V]) Keys() []K {
	keys := make([]K, 0, size(t.root))
	return inOrder(t.root, keys)
}

func (t *RBTree[K, V]) Rank(key K) int {
	return rank(t.root, key, t.compareFn)
}

func rank[K Comparable, V any](node *Node[K, V], key K, compareFn CompareFn[K]) int {
	if node == nil {
		return 0
	}
	cmp := compare(key, node.key, compareFn)
	if cmp < 0 {
		return rank(node.left, key, compareFn)
	}
	if cmp > 0 {
		return 1 + size(node.left) + rank(node.right, key, compareFn)
	}
	return size(node.left)
}

func inOrder[K Comparable, V any](node *Node[K, V], keys []K) []K {
	if node == nil {
		return keys
	}
	keys = inOrder(node.left, keys)
	keys = append(keys, node.key)
	keys = inOrder(node.right, keys)
	return keys
}

func isRed[K Comparable, V any](node *Node[K, V]) bool {
	if node == nil {
		return false
	}
	return node.color == Red
}

func height[K Comparable, V any](node *Node[K, V]) float64 {
	if node == nil {
		return -1
	}
	return 1 + math.Max(height(node.left), height(node.right))
}

func deleteNode[K Comparable, V any](node *Node[K, V], key K, compareFn CompareFn[K]) *Node[K, V] {
	if compare(key, node.key, compareFn) < 0 {
		if !isRed(node.left) && !isRed(node.left.left) {
			node = moveRedLeft(node)
		}
		node.left = deleteNode(node.left, key, compareFn)
		return balance(node)
	}
	if isRed(node.left) {
		node = rotateRight(node)
	}
	if compare(key, node.key, compareFn) == 0 && node.right == nil {
		return nil
	}
	if !isRed(node.right) && !isRed(node.right.left) {
		node = moveRedRight(node)
	}
	if compare(key, node.key, compareFn) == 0 {
		x := min(node.right)
		node.key = x.key
		node.val = x.val
		node.right = deleteMin(node.right)
	} else {
		node.right = deleteNode(node.right, key, compareFn)
	}
	return balance(node)
}

func deleteMax[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	if isRed(node.left) {
		node = rotateRight(node)
	}
	if node.right == nil {
		return nil
	}
	if !isRed(node.right) && !isRed(node.right.left) {
		node = moveRedRight(node)
	}
	node.right = deleteMax(node.right)
	return balance(node)
}

func deleteMin[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	if node.left == nil {
		return nil
	}
	if !isRed(node.left) && !isRed(node.left.left) {
		node = moveRedLeft(node)
	}
	node.left = deleteMin(node.left)
	return balance(node)
}

func max[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	if node.right == nil {
		return node
	}
	return max(node.right)
}

func min[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	if node.left == nil {
		return node
	}
	return min(node.left)

}

func get[K Comparable, V any](node *Node[K, V], key K, compareFn CompareFn[K]) (V, bool) {
	if node == nil {
		return (&Node[K, V]{}).val, false // hack to workaround generics zero value with correct type
	}
	cmp := compare(key, node.key, compareFn)
	if cmp < 0 {
		return get(node.left, key, compareFn)
	}
	if cmp > 0 {
		return get(node.right, key, compareFn)
	}
	return node.val, true
}

func put[K Comparable, V any](node *Node[K, V], key K, val V, compareFn CompareFn[K]) *Node[K, V] {
	if node == nil {
		return &Node[K, V]{
			key:   key,
			val:   val,
			size:  1,
			color: Red,
		}
	}

	cmp := compare(key, node.key, compareFn)
	if cmp < 0 {
		node.left = put(node.left, key, val, compareFn)
	} else if cmp > 0 {
		node.right = put(node.right, key, val, compareFn)
	} else {
		node.val = val
	}
	return balance(node)

}

func size[K Comparable, V any](node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return int(node.size)
}

func compare[K Comparable](key1, key2 K, compareFn CompareFn[K]) int {
	if compareFn != nil {
		return compareFn(key1, key2)
	}
	if key1 < key2 {
		return -1
	}
	if key1 > key2 {
		return 1
	}
	return 0
}

func rotateLeft[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = Red
	x.size = node.size
	node.size = size(node.left) + size(node.right) + 1
	return x

}

func rotateRight[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = Red
	x.size = node.size
	node.size = size(node.left) + size(node.right) + 1
	return x
}

func flipColors[K Comparable, V any](node *Node[K, V]) {
	// bitwise xor to flip colors (0 and 1)
	node.color ^= node.color
	node.left.color ^= node.left.color
	node.right.color ^= node.right.color
}

func balance[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}
	node.size = size(node.left) + size(node.right) + 1
	return node
}

func moveRedRight[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	flipColors(node)
	if isRed(node.left.left) {
		node = rotateRight(node)
		flipColors(node)
	}
	return node
}

func moveRedLeft[K Comparable, V any](node *Node[K, V]) *Node[K, V] {
	flipColors(node)
	if isRed(node.right.left) {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
		flipColors(node)
	}
	return node
}
