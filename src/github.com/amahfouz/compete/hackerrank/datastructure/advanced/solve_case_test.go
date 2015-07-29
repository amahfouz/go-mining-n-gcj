package advanced

import "testing"

// import "fmt"
import "github.com/amahfouz/util/test"

func TestSingleNode(t *testing.T) {
	result := solve([]int64{1}, []int64{80})
	test.Assert0(t, result == 80)
}

func TestAscending(t *testing.T) {
	result := solve([]int64{1, 2, 3, 4}, []int64{70, 20, 10, 5})
	test.Assert0(t, result == 105)
}

func TestVarious(t *testing.T) {
	result := solve([]int64{1, 2, 4, 3}, []int64{70, 20, 10, 5})
	test.Assert0(t, result == 100)

	result = solve([]int64{1, 3, 4, 2, 3}, []int64{40, 20, 10, 30, 10})
	test.Assert0(t, result == 80)
}

func TestHackerRankExample(t *testing.T) {
	result := solve([]int64{1, 2, 3, 4, 1, 2, 3, 4}, []int64{10, 20, 30, 40, 15, 15, 15, 50})
	test.Assert0(t, result == 110)
}
