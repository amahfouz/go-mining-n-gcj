package test

import "testing"
import "reflect"

// marker interface
type Object interface{}

func Assert(t *testing.T, condition bool, errMsg string) {
	if !condition {
		t.Error(errMsg)
	}
}

func AssertEquals(t *testing.T, x Object, y Object) {
	if x != y {
		t.Error("Values expected to be equal.")
	}
}

func AssertNotEquals(t *testing.T, x Object, y Object) {
	if x == y {
		t.Error("Values expected to be different.")
	}
}

func Assert0(t *testing.T, condition bool) {
	Assert(t, condition, "Assertion failed.")
}

func AssertNonNil(t *testing.T, nonNullObj Object) {
	if nonNullObj == nil {
		t.Error("Unexpected null.")
	}
}

func AssertNil(t *testing.T, nullObj Object) {
	if !reflect.ValueOf(nullObj).IsNil() {
		t.Error(nullObj)
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
