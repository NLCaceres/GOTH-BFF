package test

import (
	"errors"
	"testing"
)

func TestErrorMsg(t *testing.T) {
	tests := map[string]struct {
		Name      string
		ExpectArg any
		ActualArg any
		Expect    string
	}{
		"String Arg": {"Foo", "Bar", "Fizz", "Expected Foo = Bar but got Fizz"},
		"Bool Arg":   {"bool", true, false, "Expected bool = true but got false"},
		"Slice Arg":  {"a", []int{1, 2, 3}, []int{1, 2}, "Expected a = [1 2 3] but got [1 2]"},
		"Map Arg": { // Doesn't actual do any comparisons so matching values don't matter
			"m", map[int]bool{1: false, 2: true}, map[int]bool{1: false, 2: true},
			"Expected m = map[1:false 2:true] but got map[1:false 2:true]",
		},
		"Error arg": {
			"error", errors.New("foo"), errors.New("bar"), "Expected error = foo but got bar",
		},
		"Struct arg": {
			"", struct{ A, B int }{1, 2}, struct{ A, B int }{1, 2}, "Expected = {1 2} but got {1 2}",
		},
		"Pointer arg": {
			"", &struct{ A int }{1}, &struct{ A int }{1}, "Expected = &{1} but got &{1}",
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := ErrorMsg(testCase.Name, testCase.ExpectArg, testCase.ActualArg)
			if actual != testCase.Expect {
				t.Error(ErrorMsg("error message", testCase.Expect, actual))
			}
		})
	}
}
