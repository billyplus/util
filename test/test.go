package test

import (
	"runtime/debug"
	"testing"
)

// T wrap testing.T for addtion function
type T testing.T

// Assert check the condition if false then error out message
func (t *T) Assert(ok bool, format string, args ...interface{}) {
	if !ok {
		t.Errorf(format, args...)
	}
}

// AssertNil check errors
// if not nil then error out message and stack
func (t *T) AssertNil(err error) {
	if err != nil {
		t.Error(t)
		t.Log(string(debug.Stack()))
	}
}
