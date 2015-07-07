package test

import "testing"

// marker interface
type NonNillable interface{}

func Assert(t *testing.T, condition bool, errMsg string) {
	if !condition {
		t.Error(errMsg)
	}
}

func AssertNonNil(t *testing.T, nonNullObj NonNillable) {
	if nonNullObj == nil {
		t.Error("Unexpected null.")
	}
}

func AssertFail(t *testing.T, fn func(), errMsg string) {
	defer func() {
		if r := recover(); r != nil {
			// failure occured as expected
		} else {
			t.Error(errMsg)
		}
	}()

	// call the function that is expected to fail
	fn()
}
