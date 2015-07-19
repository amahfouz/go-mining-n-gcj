package collection

import "testing"
import "github.com/amahfouz/util/test"

type Person struct {
	age uint
}

func (p Person) SmallerThan(other Comparable) bool {
	return p.age < other.(Person).age
}

func TestComparable(t *testing.T) {
	old := Person{50}
	yng := Person{20}

	test.Assert(t, yng.SmallerThan(old), "Comparison failed.")
}

func TestEqualsComparableSlices(t *testing.T) {
	p := []Comparable{Person{10}, Person{40}, Person{20}}
	q := []Comparable{Person{10}, Person{40}, Person{20}}
	r := []Comparable{Person{10}, Person{40}}
	s := []Comparable{Person{10}, Person{40}, Person{30}}

	test.Assert0(t, EqualsSlices(p, q))
	test.Assert0(t, !EqualsSlices(p, r))
	test.Assert0(t, !EqualsSlices(s, r))
	test.Assert0(t, !EqualsSlices(s, q))
}

func TestInsertSortedDec(t *testing.T) {
	p := []Comparable{Person{40}, Person{10}}
	q := InsertSortedDec(p, Person{20})
	test.Assert(t, EqualsSlices([]Comparable{Person{40}, Person{20}, Person{10}}, q), "Middle insert.")

	p = []Comparable{Person{10}}
	q = InsertSortedDec(p, Person{20})
	test.Assert(t, EqualsSlices([]Comparable{Person{20}, Person{10}}, q), "Beginning insert.")

	p = []Comparable{Person{40}, Person{20}}
	q = InsertSortedDec(p, Person{10})
	test.Assert(t, EqualsSlices([]Comparable{Person{40}, Person{20}, Person{10}}, q), "End insert.")

	p = []Comparable{}
	q = InsertSortedDec(p, Person{10})
	test.Assert(t, EqualsSlices([]Comparable{Person{10}}, q), "Empty insert.")
}
