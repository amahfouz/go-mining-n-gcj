package test

import "testing"

func Assert(t *testing.T, condition bool, errMsg string) {
    if ! condition {
        t.Error(errMsg)
    }
}