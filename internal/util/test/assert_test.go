package test

import (
	"errors"
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"reflect"
	"testing"
)

func TestIsBothNil(t *testing.T) {
	tests := map[string]assertionTest{
		"Two nil values":        {nil, nil, true}, // Ensure two nils return true
		"Two 0s":                {0, 0, false},    // All else should be false
		"Two empty strings":     {"", "", false},
		"A 0 and empty string":  {0, "", false}, // JUST want to be sure both nil, not falsy
		"Two arrays":            {[]int{}, []int{}, false},
		"Two made empty arrays": {make([]int, 0), make([]int, 0), false},
		"Two made arrays":       {make([]int, 1), make([]int, 1), false},
		"Two maps":              {map[string]int{}, map[string]int{}, false},
		"Two made maps":         {make(map[string]int), make(map[string]int), false},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if IsBothNil(testCase.Lhs, testCase.Rhs) != testCase.Expect {
				if testCase.Lhs == nil || testCase.Rhs == nil {
					t.Error("Two nil values unexpectedly non-nil")
				} else {
					t.Errorf("Two %vs unexpectedly = %v vs %v\n", reflect.TypeOf(testCase.Lhs).Kind(), testCase.Lhs, testCase.Rhs)
				}
			}
		})
	}
}

func TestIsBothNonNil(t *testing.T) {
	tests := map[string]assertionTest{
		"Two nil values":        {nil, nil, false}, // Ensure two nils return false
		"Two 0s":                {0, 0, true},      // All else should be true
		"Two empty strings":     {"", "", true},
		"A 0 and empty string":  {0, "", true}, // JUST want to be sure both NOT nil
		"Two arrays":            {[]int{}, []int{}, true},
		"Two made empty arrays": {make([]int, 0), make([]int, 0), true},
		"Two made arrays":       {make([]int, 1), make([]int, 1), true},
		"Two maps":              {map[string]int{}, map[string]int{}, true},
		"Two made maps":         {make(map[string]int), make(map[string]int), true},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if IsBothNonNil(testCase.Lhs, testCase.Rhs) != testCase.Expect {
				if testCase.Lhs == nil || testCase.Rhs == nil {
					t.Error("Two nil values unexpectedly non-nil")
				} else {
					t.Errorf("Two %vs unexpectedly = %v vs %v\n", reflect.TypeOf(testCase.Lhs).Kind(), testCase.Lhs, testCase.Rhs)
				}
			}
		})
	}
}

func TestOnlyOneIsNil(t *testing.T) {
	tests := map[string]assertionTest{ //NOTE: Why have `OnlyOneIsNil`?
		// The first 3 cases illustrate: Return true if ONE value is nil BUT the other is non-nil
		"One nil value & empty string": {"", nil, true},   // Only one is nil, so return true
		"Two nil values":               {nil, nil, false}, // BOTH nil returns false
		"Two 0s":                       {0, 0, false},     // All else should be false too
		"Two empty strings":            {"", "", false},
		"A 0 and empty string":         {0, "", false}, // Not checking falsy equality or type
		"Two arrays":                   {[]int{}, []int{}, false},
		"Two made empty arrays":        {make([]int, 0), make([]int, 0), false},
		"Two made arrays":              {make([]int, 1), make([]int, 1), false},
		"Two maps":                     {map[string]int{}, map[string]int{}, false},
		"Two made maps":                {make(map[string]int), make(map[string]int), false},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if OnlyOneIsNil(testCase.Lhs, testCase.Rhs) != testCase.Expect {
				if testCase.Lhs == nil || testCase.Rhs == nil {
					t.Error("Two nil values unexpectedly found to have 1 non-nil value")
				} else {
					t.Errorf("Two %vs = %v & %v unexpectedly found to have 1 non-nil value\n", reflect.TypeOf(testCase.Lhs).Kind(), testCase.Lhs, testCase.Rhs)
				}
			}
		})
	}
}

// For table tests that take two values, compare them, either to each other or specific
// values, and expect the two comparisons to yield a particular logic-based bool result
type assertionTest struct {
	Lhs    any
	Rhs    any
	Expect bool
}

func TestIsSameError(t *testing.T) {
	tests := map[string]struct {
		Err    error
		Msg    string
		Expect bool
	}{
		"No error":                            {nil, "", true},
		"No error BUT expect message":         {nil, "Foo", false},
		"Error BUT not expecting it":          {errors.New("Foo"), "", false},
		"Error AND expecting it":              {errors.New("Foo"), "Foo", true},
		"Error AND expecting similar message": {errors.New("Foo Bar"), "Foo", true},
		"Error AND expecting longer message":  {errors.New("Foo"), "Foo Bar", false},
		"Error AND expecting different one":   {errors.New("Bar"), "Foo", false},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if IsSameError(testCase.Err, testCase.Msg) != testCase.Expect {
				t.Error(ErrorMsg("errors the same", testCase.Expect, !testCase.Expect))
			}
		})
	}
}

func TestEqualErrors(t *testing.T) {
	sentinel := errors.New("foobar")
	tests := map[string]struct {
		Err     error
		ErrType any
		Expect  bool
	}{
		"Nil error EQUAL to  nil":              {nil, nil, true},
		"Nil error NOT equal to an error":      {nil, errors.New("foo"), false},
		"Nil error NOT equal to custom error":  {nil, new(fileread.Error), false},
		"Same sentinel err instance NOT equal": {sentinel, sentinel, false},
		"error NOT equal to a similar error":   {errors.New("foo"), errors.New("foo"), false},
		"error NOT equal to nil":               {errors.New("foo"), nil, false},
		"error NOT equal to another error":     {errors.New("foo"), errors.New("bar"), false},
		"error NOT equal to another errorF": {
			errors.New("foo"), fmt.Errorf("%v", "foo"), false,
		},
		"error NOT equal to custom error": {
			errors.New("foo"), new(fileread.FileNotFoundError), false,
		},
		"error NOT equal to ANY custom error interface": {
			errors.New("foo"), new(fileread.Error), false,
		},
		"error NOT equal to true base any interface": {
			errors.New("foo"), new(any), false,
		},
		"error EQUAL to base error interface": {
			errors.New("foo"), new(error), true,
		},
		"errorF NOT equal to nil": {fmt.Errorf("%v", "foo"), nil, false},
		"errorF NOT equal to a similar errorF": {
			fmt.Errorf("%v", "foo"), fmt.Errorf("%v", "foo"), false,
		},
		"errorF NOT equal to another errorF": {
			fmt.Errorf("%v", "foo"), fmt.Errorf("%v", "bar"), false,
		},
		"errorF NOT equal to custom error": {
			fmt.Errorf("%v", "foo"), new(fileread.FileNotFoundError), false,
		},
		"errorF NOT equal to ANY custom error interface": {
			fmt.Errorf("%v", "foo"), new(fileread.Error), false,
		},
		"errorF NOT equal to true base any interface": {
			fmt.Errorf("%v", "foo"), new(any), false,
		},
		"errorF EQUAL to base error interface": {
			fmt.Errorf("%v", "foo"), new(error), true,
		},
		"Custom error EQUAL to implemented interface error type": {
			fileread.FileNotFoundError{}, new(fileread.Error), true,
		},
		"Custom error NOT equal to true base any interface": {
			fileread.FileNotFoundError{}, new(any), false,
		},
		"Custom error EQUAL to base error interface": {
			fileread.FileNotFoundError{}, new(error), true,
		},
		"Custom error EQUAL to same error struct pointer type": {
			fileread.FileNotFoundError{}, new(fileread.FileNotFoundError), true,
		},
		"Custom error with message EQUAL to same error struct pointer type": {
			fileread.FileNotFoundError{File: "foo"}, new(fileread.FileNotFoundError), true,
		},
		"Custom error NOT equal to same error struct type": {
			fileread.FileNotFoundError{}, fileread.FileNotFoundError{}, false,
		},
		"Custom error NOT equal to different custom error struct pointer": {
			fileread.FileNotFoundError{}, new(fileread.MalformedJsonError), false,
		},
		"Custom error NOT equal to different custom error struct": {
			fileread.FileNotFoundError{}, fileread.MalformedJsonError{}, false,
		},
		"Custom error NOT equal to nil": {fileread.FileNotFoundError{}, nil, false},
		"Custom error NOT equal to an error": {
			fileread.FileNotFoundError{}, errors.New("foo"), false,
		},
		"Custom error NOT equal to an errorF": {
			fileread.FileNotFoundError{}, fmt.Errorf("%v", "foo"), false,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if EqualErrors(testCase.Err, testCase.ErrType) != testCase.Expect {
				t.Error(ErrorMsg("Equal Errors", testCase.Expect, !testCase.Expect))
			}
		})
	}
}
