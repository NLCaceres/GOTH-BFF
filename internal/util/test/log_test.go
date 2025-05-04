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
		"Nil arg": {"", struct{ A, B int }{1, 2}, nil, "Expected = {1 2} but got <nil>"},
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

func TestQuotedErrorMsg(t *testing.T) {
	tests := map[string]struct {
		Name      string
		ExpectArg any
		ActualArg any
		Expect    string
	}{
		"String Arg": {"Foo", "Bar", "Fizz", `Expected Foo = "Bar" but got "Fizz"`},
		// Bools CAN'T be used with the %q fmt verb, so original output is altered
		"Bool Arg": {"bool", true, false, `Expected bool = "true" but got "false"`},
		"Slice Arg": { // Ints get converted to a character literal
			"a", []int{1, 2, 3}, []int{1, 2},
			"Expected a = ['\\x01' '\\x02' '\\x03'] but got ['\\x01' '\\x02']",
		},
		"Map Arg": { // Values can match since the func just outputs a string, no comparison done
			"m", map[string]string{"1": "a", "2": "b"}, map[string]string{"1": "a", "2": "b"},
			`Expected m = map["1":"a" "2":"b"] but got map["1":"a" "2":"b"]`,
		},
		"Error arg": {
			"error", errors.New("foo"), errors.New("bar"), `Expected error = "foo" but got "bar"`,
		},
		"Struct arg": {
			"", struct{ A, B string }{"A", "B"}, struct{ A, B string }{"A", "B"},
			`Expected = {"A" "B"} but got {"A" "B"}`,
		},
		"Nil arg": { // Nil values end up with similar odd formatting like bools
			"", struct{ A string }{"A"}, nil, `Expected = {"A"} but got "nil"`,
		},
		"Pointer arg": {
			"", &struct{ A string }{"A"}, &struct{ A string }{"A"}, `Expected = &{"A"} but got &{"A"}`,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := QuotedErrorMsg(testCase.Name, testCase.ExpectArg, testCase.ActualArg)
			if actual != testCase.Expect {
				t.Error(ErrorMsg("error message", testCase.Expect, actual))
			}
		})
	}
}
