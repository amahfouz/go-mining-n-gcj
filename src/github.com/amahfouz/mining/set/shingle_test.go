package set

import "testing"

func TestBasic(t *testing.T) {
    s := "abcd"
    shingles := Shingle(s, 2)
    assert(t, shingles.NumElements() == 3, "Wrong number of shingles")
}

func TestRepeating(t *testing.T) {
    s := "abcabd"
    shingles := Shingle(s, 2)
    assert(t, shingles.NumElements() == 4, "Wrong number of shingles")
}

func TestOneShingle(t *testing.T) {
    s := "abcd"
    shingles := Shingle(s, 4)
    assert(t, shingles.NumElements() == 1, "Wrong number of shingles")
}

func assert(t *testing.T, condition bool, errMsg string) {
    if ! condition {
        t.Error(errMsg)
    }
}