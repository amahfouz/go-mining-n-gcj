package collection

import "testing"
import "github.com/amahfouz/util/test"

func TestBuildTree(t *testing.T) {
	var data = []int64{1, 3, 5}
	tree := NewSegTree(data)
	test.AssertNonNil(t, tree)

	test.Assert(t, EqualsInt64(tree.data, []int64{1, 3, 5}), "Tree data.")
	test.Assert(t, EqualsInt(tree.nodes, []int{-1, -1, 2, 0, 1}), "Tree nodes.")
	test.Assert(t, EqualsInt(tree.lo, []int{0, 0, 2, 0, 1}), "Tree low.")
	test.Assert(t, EqualsInt(tree.hi, []int{2, 1, 2, 0, 1}), "Tree high.")
}
