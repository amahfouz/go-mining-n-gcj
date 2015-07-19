package collection

type Int64arr []int64

type Comparable interface {
	SmallerThan(Comparable) bool
}

// make type Int64arr sortable

func (a Int64arr) Len() int {
	return len(a)
}

func (a Int64arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Int64arr) Less(i, j int) bool {
	return a[i] < a[j]
}

// Insertion sort

func InsertSortedDec(s []Comparable, a Comparable) []Comparable {
	index := 0
	size := len(s)

	for index < size && a.SmallerThan(s[index]) {
		index++
	}

	var result []Comparable
	if index < size {
		if cap(s) == size {
			result = append(s, nil)
		}
		copy(result[index+1:], result[index:])
		result[index] = a
	} else {
		result = append(s, a)
	}
	return result
}

// Slice of Comparable

func EqualsSlices(a, b []Comparable) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
