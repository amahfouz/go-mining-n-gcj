package advanced

import (
	//	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

import "github.com/amahfouz/util/test"

func TestConstruct(t *testing.T) {
	tree := NewTreap()
	test.AssertNil(t, tree.root)
}

func TestAddTreapRootNode(t *testing.T) {
	tree := createTreap()
	n := createNode(5)

	tree.Insert(n)

	expected := createNode(5)

	test.Assert0(t, DeepEquals(tree.root, expected))
}

func TestAddTreapLeftAndRightNodes(t *testing.T) {
	tree := createTreap()
	n := createNode(50)
	n.totalWeight = 80
	m := createNode(50)
	m.totalWeight = 60
	tree.Insert(n)
	tree.Insert(m)

	expectedRoot := createNode(50)
	expectedRoot.totalWeight = 80

	expectedLeft := createNode(50)
	expectedLeft.totalWeight = 60

	expectedLeft.parent = expectedRoot
	expectedRoot.left = expectedLeft

	test.Assert0(t, DeepEquals(tree.root, expectedRoot))

	r := createNode(50)
	r.totalWeight = 90
	tree.Insert(r)

	expectedRight := createNode(50)
	expectedRight.totalWeight = 90

	expectedRight.parent = tree.root
	expectedRoot.right = expectedRight

	test.Assert0(t, DeepEquals(tree.root, expectedRoot))
}

func TestAddTreapLeftRightNodes(t *testing.T) {
	tree := createTreap()
	tree.Insert(createNodeWithWeight(100))
	tree.Insert(createNodeWithWeight(50))

	expectedLeft := createNodeWithWeight(50)
	expectedLeft.parent = tree.root
	expectedRoot := createNodeWithWeight(100)
	expectedRoot.left = expectedLeft

	tree.Insert(createNodeWithWeight(80))
	expectedRight := createNodeWithWeight(80)
	expectedRight.parent = expectedLeft
	expectedLeft.right = expectedRight

	test.Assert0(t, DeepEquals(tree.root, expectedRoot))
}

func TestRotateLeftThenRight(t *testing.T) {
	tree := createTreapWithPriorities([]int64{3, 4, 1})
	tree.Insert(createNodeWithWeight(10))
	tree.Insert(createNodeWithWeight(6))
	tree.Insert(createNodeWithWeight(8))

	tree2 := createTreap()
	tree2.Insert(createNodeWithWeight(8))
	tree2.Insert(createNodeWithWeight(6))
	tree2.Insert(createNodeWithWeight(10))

	test.Assert0(t, DeepEquals(tree.root, tree2.root))
}

func TestSingleRotateLeft(t *testing.T) {
	tree := createTreapWithPriorities([]int64{3, 5, 4})
	tree.Insert(createNodeWithWeight(10))
	tree.Insert(createNodeWithWeight(6))
	tree.Insert(createNodeWithWeight(8))

	tree2 := createTreap()
	tree2.Insert(createNodeWithWeight(10))
	tree2.Insert(createNodeWithWeight(8))
	tree2.Insert(createNodeWithWeight(6))

	test.Assert0(t, DeepEquals(tree.root, tree2.root))
}

func TestSingleRotateRight(t *testing.T) {
	tree := createTreapWithPriorities([]int64{3, 5, 4})
	tree.Insert(createNodeWithWeight(6))
	tree.Insert(createNodeWithWeight(8))
	tree.Insert(createNodeWithWeight(7))

	tree2 := createTreap()
	tree2.Insert(createNodeWithWeight(6))
	tree2.Insert(createNodeWithWeight(7))
	tree2.Insert(createNodeWithWeight(8))

	test.Assert0(t, DeepEquals(tree.root, tree2.root))
}

func TestMultipleAddRebalance(t *testing.T) {
	tree := NewTreap()
	var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

	var i int64
	for i = 0; i < 1000; i++ {
		n := createNode(randomizer.Int63n(math.MaxInt64) % 100000)
		n.index = i + 1
		tree.Insert(n)
	}
	test.Assert(t, tree.isValid(tree.root), "Tree semantics not preserved.")
}

// helpers

func createTreap() *Treap {
	genFixed := func() int64 {
		return 0
	}

	tree := new(Treap)
	tree.nextRandomPriority = genFixed

	return tree
}

func createTreapWithPriorities(priorities []int64) *Treap {
	i := -1
	genRand := func() int64 {
		i++
		return priorities[i]
	}

	tree := new(Treap)
	tree.nextRandomPriority = genRand

	return tree
}

func createNode(data int64) *Node {
	n := new(Node)
	n.data = data
	n.lo = data

	return n
}

func createNodeWithWeight(totalWeight int64) *Node {
	n := new(Node)
	n.totalWeight = totalWeight
	return n
}
