package collection

import (
	"fmt"
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
	tree.Add(1, 100)

	expected := Node{1, 100, 0, nil, nil, nil}

	test.Assert0(t, DeepEquals(tree.root, &expected))
}

func TestAddTreapLeftAndRightNodes(t *testing.T) {
	tree := createTreap()
	tree.Add(1, 100)
	tree.Add(3, 50)

	expectedNode := Node{3, 50, 0, nil, nil, tree.root}
	expectedRoot := Node{1, 100, 0, &expectedNode, nil, nil}
	test.Assert0(t, DeepEquals(tree.root, &expectedRoot))

	tree.Add(4, 120)
	expectedRight := Node{4, 120, 0, nil, nil, tree.root}
	expectedRoot.right = &expectedRight

	test.Assert0(t, DeepEquals(tree.root, &expectedRoot))
}

func TestAddTreapLeftRightNodes(t *testing.T) {
	fmt.Println("TestAddTreapLeftRightNodes")
	tree := createTreap()
	tree.Add(1, 100)
	tree.Add(3, 50)

	expectedLeft := Node{3, 50, 0, nil, nil, tree.root}
	expectedRoot := Node{1, 100, 0, &expectedLeft, nil, nil}

	tree.Add(4, 80)
	expectedRight := Node{4, 80, 0, nil, nil, &expectedLeft}
	expectedLeft.right = &expectedRight

	test.Assert0(t, DeepEquals(tree.root, &expectedRoot))
}

func TestRotateLeftThenRight(t *testing.T) {
	tree := createTreapWithPriorities([]int64{3, 4, 1})
	tree.Add(1, 10)
	tree.Add(2, 6)
	tree.Add(3, 8)

	tree2 := createTreap()
	tree2.Add(3, 8)
	tree2.Add(2, 6)
	tree2.Add(1, 10)

	test.Assert0(t, DeepEquals(tree.root, tree2.root))
}

func TestSingleRotateLeft(t *testing.T) {
	tree := createTreapWithPriorities([]int64{3, 5, 4})
	tree.Add(1, 10)
	tree.Add(2, 6)
	tree.Add(3, 8)

	tree2 := createTreap()
	tree2.Add(1, 10)
	tree2.Add(3, 8)
	tree2.Add(2, 6)

	test.Assert0(t, DeepEquals(tree.root, tree2.root))
}

func TestSingleRotateRight(t *testing.T) {
	tree := createTreapWithPriorities([]int64{3, 5, 4})
	tree.Add(2, 6)
	tree.Add(1, 8)
	tree.Add(3, 7)

	tree2 := createTreap()
	tree2.Add(2, 6)
	tree2.Add(3, 7)
	tree2.Add(1, 8)

	test.Assert0(t, DeepEquals(tree.root, tree2.root))
}

func TestMultipleAddRebalance(t *testing.T) {
	tree := NewTreap()
	var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

	var i int64
	for i = 0; i < 1000; i++ {
		tree.Add(i, randomizer.Int63n(math.MaxInt64))
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
